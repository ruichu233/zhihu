package logic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishContentLogic {
	return &PublishContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 当创作者发布新内容时，推送内容发布事件
func (l *PublishContentLogic) PublishContent(in *feed.PublishContentRequest) (*feed.PublishContentResponse, error) {
	l.svcCtx.RDB.ZAdd(l.ctx, GetCacheKey(in.UserId), redis.Z{
		Score:  float64(in.VideoCreatorTimestamp),
		Member: in.VideoId,
	})
	// 遍历关注列表，给每个粉丝推送新内容事件
	// 1、获取粉丝列表
	followerList, err := l.svcCtx.FollowRPC.ListFollowers(l.ctx, &follow.GetFollowerListRequest{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	// 2、推送新内容事件
	for _, followerId := range followerList.FollowerIds {
		// 3、将新内容事件推送到粉丝的缓存列表中
		intCmd := l.svcCtx.RDB.ZAdd(l.ctx, GetCacheKey(followerId), redis.Z{
			Score:  float64(in.VideoCreatorTimestamp),
			Member: in.VideoId,
		})
		if intCmd.Val() <= 0 {
			logx.Errorf("用户 %d 推送新内容失败", followerId)
		}
	}

	return &feed.PublishContentResponse{
		Success: true,
	}, nil
}

func GetCacheKey(userId int64) string {
	return fmt.Sprintf("FEED_CONTENT_%d", userId)
}
