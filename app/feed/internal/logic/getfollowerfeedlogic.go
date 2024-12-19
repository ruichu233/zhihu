package logic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"

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
	// 1、缓存
	key := GetCacheKey(in.UserId)
	result, err := l.svcCtx.RDB.ZRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Count:  10,
		Offset: 0,
	}).Result()
	if err != nil {
		return nil, err
	}
	itemIds := make([]int64, 0, len(result))
	for _, v := range result {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		itemIds = append(itemIds, id)
	}
	return &feed.GetFollowerFeedResponse{
		Items:      itemIds,
		NextCursor: 0,
		HasMore:    false,
	}, nil
}
