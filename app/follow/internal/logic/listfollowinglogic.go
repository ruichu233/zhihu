package logic

import (
	"context"
	"encoding/json"
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

// ListFollowing 获取关注列表
func (l *ListFollowingLogic) ListFollowing(in *follow.GetFollowListRequest) (*follow.GetFollowListResponse, error) {
	if in.PageSize == 0 {
		in.PageSize = 10
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}
	var (
		isCache, isEnd bool
		cursor, lastId int64
		followIdList   []int64
		follows        []*model.Follow
		curPage        []*follow.FollowItem
	)

	// 1、构建获取关注id列表缓存键
	followKey := GetFollowingKey(in.UserId)
	// 2、从缓存中获取关注id列表
	followIdList, _ = l.cacheFollowing(followKey, in.Cursor, in.PageSize)
	if len(followIdList) > 0 {
		isCache = true
		if followIdList[len(followIdList)-1] == -1 {
			followIdList = followIdList[:len(followIdList)-1]
			isEnd = true
		}
		if len(followIdList) == 0 {
			return &follow.GetFollowListResponse{}, nil
		}
		_follows, err := l.getFollowingsListByIds(in.UserId, followIdList)
		if err != nil {
			return nil, err
		}
		follows = _follows
		for _, v := range follows {
			curPage = append(curPage, &follow.FollowItem{
				Id:         v.Id,
				UserId:     v.FolloweeID,
				CreateTime: v.CreatedAt.Unix(),
			})
		}
	} else {
		defaultCache := 10 * in.PageSize
		createTime := time.Unix(in.Cursor, 0)
		if err := l.svcCtx.DB.Model(&model.Follow{}).
			Where("follower_id = ? AND updated_at <?", in.UserId, createTime).
			Limit(int(defaultCache)).
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
			followIdList = append(followIdList, f.FolloweeID)
			curPage = append(curPage, &follow.FollowItem{
				Id:         f.Id,
				UserId:     f.FolloweeID,
				CreateTime: f.CreatedAt.Unix(),
			})
		}
	}
	// 根据 Id 去重
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		cursor = pageLast.CreateTime
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
			if len(follows) < int(10*in.PageSize) && len(follows) > 0 {
				follows = append(follows, &model.Follow{
					FolloweeID: -1,
				})
				for _, v := range follows {
					l.svcCtx.RDB.ZAdd(l.ctx, followKey, redis.Z{
						Score:  float64(v.UpdatedAt.Unix()),
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
	maxScore := "inf"
	if cursor > 0 {
		maxScore = fmt.Sprintf("%d", cursor)
	}
	result, err := l.svcCtx.RDB.ZRevRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    maxScore,
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

// 根据id批量查询关注信息
func (l *ListFollowingLogic) getFollowingsListByIds(userId int64, followingIds []int64) ([]*model.Follow, error) {
	key := fmt.Sprintf("follow_model:%d", userId)
	mkeys := make([]string, 0, len(followingIds))
	for _, v := range followingIds {
		mkeys = append(mkeys, fmt.Sprintf("%d", v))
	}
	result, _ := l.svcCtx.RDB.HMGet(l.ctx, key, mkeys...).Result()
	followSet := make(map[int64]*model.Follow)
	for _, v := range result {
		if v == nil {
			continue
		}
		if v == "" {
			continue
		}
		_follow := &model.Follow{}
		err := json.Unmarshal([]byte(v.(string)), _follow)
		if err != nil {
			return nil, err
		}
		followSet[_follow.FolloweeID] = _follow
	}
	noHit := make([]int64, 0)
	for _, v := range followingIds {
		if _, ok := followSet[v]; !ok {
			noHit = append(noHit, v)
		}
	}
	// 如果缓存中的获取的数据不够
	if len(noHit) > 0 {
		var list []*model.Follow
		// 在数据库中查询缓存没有的数据，并且存入缓存
		err := l.svcCtx.DB.Model(&model.Follow{}).
			Where("followee_id in ? and follower_id = ?", noHit, userId).
			Find(&list).Error
		if err != nil {
			return nil, err
		}
		for _, v := range list {
			followSet[v.FolloweeID] = v
			var b []byte
			if b, err = json.Marshal(v); err != nil {
				return nil, err
			}
			_ = l.svcCtx.RDB.HSet(l.ctx, key, fmt.Sprintf("%d", v.FolloweeID), string(b)).Err()
		}
	}
	// 最后组装
	var resultList []*model.Follow
	for _, v := range followingIds {
		if _follow, ok := followSet[v]; ok {
			resultList = append(resultList, _follow)
		}
	}
	return resultList, nil
}
