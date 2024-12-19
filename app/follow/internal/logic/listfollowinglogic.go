package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/model"
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
	var (
		isCache, isEnd bool
		cursor, lastId int64
		follows        []*model.Follow
		curPage        []*follow.FollowItem
	)

	// 构建缓存键
	followKey := GetFollowKey(in.UserId)
	ids, _ := l.cacheFollowing(followKey, in.Cursor, in.PageSize)
	if len(ids) > 0 {
		isCache = true
		if ids[len(ids)-1] == -1 {
			ids = ids[:len(ids)-1]
			isEnd = true
		}
		if len(ids) == 0 {
			return &follow.GetFollowListResponse{}, nil
		}
		var follows []*model.Follow
		if err := l.svcCtx.DB.Model(&model.Follow{}).Where("follower_id = ?").Where("followee_id IN ?", ids).Find(&follows).Error; err != nil {
			return nil, err
		}
		for _, v := range follows {
			curPage = append(curPage, &follow.FollowItem{
				Id:         v.Id,
				FolloweeId: v.FolloweeID,
				CreateTime: v.CreatedAt.Unix(),
			})
		}
	} else {
		var follows []*model.Follow
		cursorToTime := time.Unix(in.Cursor, 0)
		if err := l.svcCtx.DB.Model(&model.Follow{}).
			Where("follower_id = ? AND updated_at <?", in.UserId, cursorToTime).
			Limit(1000).
			Find(&follows).Error; err != nil {
			return nil, err
		}
		if len(follows) == 0 {
			return &follow.GetFollowListResponse{}, nil
		}
		var firstPageFollows []*model.Follow
		if len(follows) > int(in.PageSize) {
			firstPageFollows = follows[:in.PageSize]
		} else {
			firstPageFollows = follows
			isEnd = true
		}
		for _, f := range firstPageFollows {
			curPage = append(curPage, &follow.FollowItem{
				Id:         f.Id,
				FolloweeId: f.FolloweeID,
				CreateTime: f.CreatedAt.Unix(),
			})
		}
	}
	// 根据 Id 去重
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		for k, v := range curPage {
			if in.Cursor == v.CreateTime && v.Id == in.Id {
				curPage = curPage[k:]
				break
			}
		}
	}
	// 构建缓存
	if !isCache {
		go func() {
			if len(follows) < 1000 && len(follows) > 0 {
				follows = append(follows, &model.Follow{
					BaseModel: model.BaseModel{
						Id: -1,
					},
				})
				for _, v := range follows {
					var score int64
					if v.Id == -1 {
						score = 0
					} else {
						score = v.UpdatedAt.Unix()
					}
					l.svcCtx.RDB.ZAdd(l.ctx, followKey, redis.Z{
						Score:  float64(score),
						Member: v.FolloweeID,
					})
				}
			}

		}()
	}

	return &follow.GetFollowListResponse{
		Cursor: cursor,
		IsEnd:  isEnd,
		Items:  curPage,
		LastId: lastId,
	}, nil
}

func (l *ListFollowingLogic) cacheFollowing(key string, cursor, ps int64) ([]int64, error) {
	res, err := l.svcCtx.RDB.Exists(l.ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return []int64{}, nil
		}
		return nil, err
	}
	if res > 0 {
		if err := l.svcCtx.RDB.Expire(l.ctx, key, time.Hour*24).Err(); err != nil {
			return nil, err
		}
	}
	result, err := l.svcCtx.RDB.ZRevRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "0",
		Max:    fmt.Sprintf("%d", cursor),
		Count:  ps,
		Offset: 0,
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
