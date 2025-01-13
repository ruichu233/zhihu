package logic

import (
	"context"
	"errors"
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
	// 遍历关注列表，给每个粉丝推送新内容事件
	// 1、获取粉丝列表
	followerList, err := l.svcCtx.FollowRPC.ListFollowers(l.ctx, &follow.GetFollowerListRequest{
		UserId:   in.UserId,
		Cursor:   0,
		PageSize: 0,
		Id:       0,
	})
	if err != nil {
		return nil, err
	}
	// 2、推送新内容事件
	for _, item := range followerList.Items {
		// 3、将新内容事件推送到粉丝的缓存列表中
		intCmd := l.svcCtx.RDB.ZAdd(l.ctx, GetCacheKey(item.Id), redis.Z{
			Score:  float64(in.VideoCreatorTimestamp),
			Member: in.VideoId,
		})
		if intCmd.Val() <= 0 {
			logx.Errorf("用户 %d 推送新内容失败", item.Id)
		}
	}
	return &feed.PublishContentResponse{
		Success: true,
	}, nil
}

func GetCacheKey(userId int64) string {
	return fmt.Sprintf("FEED_CONTENT_%d", userId)
}

// 检查是否还有更多数据
func checkNoMoreFlag(ctx context.Context, rdb *redis.Client, sortedSetKey string) (bool, error) {
	// 检查 `NoMore` 标志是否在 SortedSet 中
	noMoreExists, err := rdb.ZScore(ctx, sortedSetKey, "NoMore").Result()
	if err != nil {
		// 如果返回 redis.Nil，表示 NoMore 标志不存在，返回 false
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return noMoreExists > 0, nil
}
