package logic

import (
	"context"

	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerFeedLogic {
	return &GetFollowerFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注者的 Feed
func (l *GetFollowerFeedLogic) GetFollowerFeed(in *feed.GetFollowerFeedRequest) (*feed.GetFollowerFeedResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.GetFollowerFeedResponse{}, nil
}
