package logic

import (
	"context"
	"zhihu/app/video/videoclient"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishHandlerRequest, userId int64) (resp *types.PublishHandlerResponse, err error) {
	publishResponse, err := l.svcCtx.VideoRPC.Publish(l.ctx, &videoclient.PublishRequest{
		AuthorId:    userId,
		CoverUrl:    req.CoverUrl,
		Description: req.Description,
		Title:       req.Title,
		VideoUrl:    req.VideoUrl,
	})
	if err != nil {
		return nil, err
	}
	resp.VideoId = publishResponse.VideoId
	return
}
