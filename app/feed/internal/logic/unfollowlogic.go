package logic

import (
	"context"
	"zhihu/app/video/pb/video"

	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Unfollow 取消关注
func (l *UnfollowLogic) Unfollow(in *feed.UnfollowRequest) (*feed.UnfollowResponse, error) {
	// 1、获取关注者的所有作品的视频id
	videoResp, err := l.svcCtx.VideoRPC.WorkList(l.ctx, &video.WorkListRequest{UserId: in.UserId})
	if err != nil {
		return nil, err
	}
	// 2、从用户收件箱中删除关注者的所有作品视频id
	pipeline := l.svcCtx.RDB.Pipeline()
	for _, videoFeed := range videoResp.VideoFeeds {
		pipeline.ZRem(l.ctx, GetCacheKey(in.CreatorId), videoFeed.VideoId)
	}
	_, err = pipeline.Exec(l.ctx)
	if err != nil {
		return nil, err
	}
	return &feed.UnfollowResponse{}, nil
}
