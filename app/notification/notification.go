package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"zhihu/app/notification/internal/config"
	"zhihu/app/notification/internal/server"
	"zhihu/app/notification/internal/svc"
	localqueue "zhihu/app/notification/local_queue"
	"zhihu/app/notification/model"
	"zhihu/app/notification/pb/notification"
	"zhihu/pkg/idgenerator"
	"zhihu/pkg/mq"
)

var configFile = flag.String("f", "etc/notification.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	localqueue.Init(ctx.RDB, ctx.DB)
	idgenerator.InitIdGenerator(c.WorkId)
	go func() {
		ctx.Consumer.Run(func(msg *mq.MsgEntity) error {
			type notify struct {
				ToUserId   int64  `json:"to_user_id"`
				FromUserId int64  `json:"from_user_id"`
				Type       int32  `json:"type"`
				Content    string `json:"content"`
			}
			var n notify
			if err := json.Unmarshal([]byte(msg.Val), &n); err != nil {
				return err
			}

			notifyModel := &model.Notification{
				BaseModel: model.BaseModel{
					ID: idgen.NextId(),
				},
				ToUserId:   n.ToUserId,
				FromUserId: n.FromUserId,
				Type:       n.Type,
				Content:    n.Content,
			}
			localqueue.NotifyQueue.Push(notifyModel)
			return nil
		})
	}()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		notification.RegisterNotificationServiceServer(grpcServer, server.NewNotificationServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting notification rpc server at %s...\n", c.ListenOn)
	s.Start()
}
