package main

import (
	"flag"
	"fmt"
	cronjob "zhihu/app/like/cron_job"
	"zhihu/app/like/internal/config"
	"zhihu/app/like/internal/server"
	"zhihu/app/like/internal/svc"
	"zhihu/app/like/local_queue"
	"zhihu/app/like/pb/like"
	"zhihu/pkg/idgenerator"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 处理点赞消息
	local_queue.InitQueue(ctx.RDB, ctx.DB)
	defer local_queue.CloseQueue()
	// 定时任务
	cronjob.Init(ctx.RDB, ctx.DB)

	// ID生成器初始化
	idgenerator.InitIdGenerator(c.WorkId)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		like.RegisterLikeServer(grpcServer, server.NewLikeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
