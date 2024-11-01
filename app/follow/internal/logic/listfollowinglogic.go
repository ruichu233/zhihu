package logic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowingLogic {
	return &ListFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListFollowingLogic) ListFollowing(in *follow.GetFollowListRequest) (*follow.GetFollowListResponse, error) {
	if in.PageSize < 0 {
		in.PageSize = 10
	}
	if in.Cursor < 0 {
		in.Cursor = 0
	}
	// 构建缓存键
	followKey := GetFollowKey(in.UserId)
	ids, _ := l.cacheFollowing(followKey, in.Cursor, in.PageSize)

	return &follow.GetFollowListResponse{}, nil
}

func (l *ListFollowingLogic) cacheFollowing(key string, cursor, ps int64) ([]int64, error) {
	result, err := l.svcCtx.RDB.ZRevRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Count:  ps,
		Offset: cursor,
	}).Result()
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(result))
	for _, v := range result {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
