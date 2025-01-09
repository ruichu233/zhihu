package logic

import (
	"context"
	"zhihu/app/feed/pb/feed"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListRequest, userId int64) (resp *types.VideoListResponse, err error) {
	// 1、获取feed列表
	feedList := []int64{}
	if req.FeedType == 1 {
		resp, err := l.svcCtx.FeedRPC.GetFollowerFeed(l.ctx, &feed.GetFollowerFeedRequest{
			UserId:   userId,
			Cursor:   req.Cursor,
			PageSize: req.Page,
		})
		if err != nil {
			return nil, err
		}
		feedList = resp.Items
	}
	if req.FeedType == 2 {
		resp, err := l.svcCtx.FeedRPC.GetRecommendedFeed(l.ctx, &feed.GetRecommendedFeedRequest{
			UserId:   userId,
			Cursor:   req.Cursor,
			PageSize: req.Page,
		})
		if err != nil {
			return nil, err
		}
		feedList = resp.RecommendedItems
	}
	// 2、根据feed列表获取视频信息
	if len(feedList) > 0 {
		l.svcCtx.VideoRPC
	}

	// 3、根据视频信息中的作者id获取作者信息
	return
}
