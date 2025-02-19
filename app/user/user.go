package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"
	"zhihu/app/user/model"
	"zhihu/pkg/interceptors"
	"zhihu/pkg/mq"

	"zhihu/app/user/internal/config"
	"zhihu/app/user/internal/server"
	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"
	"zhihu/pkg/idgenerator"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	// 初始化本地队列
	InitQueue(ctx.RDB, ctx.DB)
	defer CloseQueue()
	// 处理用户的关注数和粉丝数，采用 先更新缓存后异步写入数据库
	ctx.Consumer.Run(func(msg *mq.MsgEntity) error {
		var action struct {
			Action uint8 `json:"action"` //1、关注数 2、粉丝数
			UserId int64 `json:"userId"`
			Type   uint8 `json:"type"` // 1、增加 2、减少
		}
		if err := json.Unmarshal([]byte(msg.Val), &action); err != nil {
			return err
		}
		// 根据 action.Action 和 action.Type 更新数据库
		var user model.User
		// 查询用户缓存
		res, err := ctx.RDB.Get(context.Background(), model.GetUserInfoKey(action.UserId)).Result()
		if err != nil && err != redis.Nil {
			return fmt.Errorf("获取缓存失败: %v", err)
		}

		// 解析缓存数据
		if res != "" {
			if res == "null" {
				return nil // 直接返回，避免查询数据库
			}
			if err := json.Unmarshal([]byte(res), &user); err != nil {
				return fmt.Errorf("解析缓存数据失败: %v", err)
			}
		} else {
			// 缓存不存在，从数据库查询
			if err := ctx.DB.Where("id = ?", action.UserId).First(&user).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// 用户不存在，设置空值缓存防止缓存穿透
					ctx.RDB.Set(context.Background(), model.GetUserInfoKey(action.UserId), "null", time.Hour*24)
					return nil
				}
				return fmt.Errorf("查询数据库失败: %v", err)
			}
			// 更新缓存
			userJson, _ := json.Marshal(user)
			ctx.RDB.Set(context.Background(), model.GetUserInfoKey(action.UserId), string(userJson), time.Hour*24)
		}

		switch action.Action {
		case 1: // 关注数
			switch action.Type {
			case 1: // 增加
				user.FollowCount++
			case 2: // 减少
				user.FollowCount--
			}
		case 2: // 粉丝数
			switch action.Type {
			case 1: // 增加
				user.FollowerCount++
			case 2: // 减少
				user.FollowerCount--
			}
		}

		// 投入本地队列处理
		UserQueue.Push(user.Id)
		return nil
	})

	// ID生成器初始化
	idgenerator.InitIdGenerator(c.WorkId)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	defer s.Stop()
	defer ctx.Consumer.Close()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
