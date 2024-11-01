package logic

import (
	"context"

	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendedFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecommendedFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendedFeedLogic {
	return &GetRecommendedFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取个性化推荐 Feed
func (l *GetRecommendedFeedLogic) GetRecommendedFeed(in *feed.GetRecommendedFeedRequest) (*feed.GetRecommendedFeedResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.GetRecommendedFeedResponse{}, nil
}
