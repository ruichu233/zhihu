package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"zhihu/app/follow/model"

	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowersLogic {
	return &ListFollowersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListFollowers 获取用户粉丝列表
func (l *ListFollowersLogic) ListFollowers(in *follow.GetFollowerListRequest) (*follow.GetFollowerListResponse, error) {
	if in.UserId == 0 {
		return nil, fmt.Errorf("user_id is empty")
	}
	if in.PageSize == 0 {
		in.PageSize = 10
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}
	var (
		isCache, isEnd  bool
		lastId, cursor  int64
		followersIdList []int64
		followers       []*model.Follow
		curPage         []*follow.FollowItem
	)
	// 构建缓存键
	key := GetFollowerKey(in.UserId)
	// 查询缓存
	followersIdList, _ = l.cacheFollowers(key, in.Cursor, in.PageSize)
	if len(followersIdList) > 0 {
		isCache = true
		if followersIdList[len(followersIdList)-1] == -1 {
			followersIdList = followersIdList[:len(followersIdList)-1]
			isEnd = true
		}
		if len(followersIdList) == 0 {
			return &follow.GetFollowerListResponse{}, nil
		}
		// 根据id批量查询关注信息
		_followers, err := l.getFollowersListByIds(in.UserId, followersIdList)
		if err != nil {
			return nil, err
		}
		followers = _followers
		for _, v := range followers {
			curPage = append(curPage, &follow.FollowItem{
				Id:         v.Id,
				UserId:     v.FollowerID,
				CreateTime: v.CreatedAt.Unix(),
			})
		}
	} else {
		// 查询数据库
		createTime := time.Unix(in.Cursor, 0)
		cacheLen := int(in.PageSize) * 10
		if err := l.svcCtx.DB.Model(&model.Follow{}).
			Where("followee_id = ? and created_at <= ?", in.UserId, createTime).Limit(cacheLen).Find(&followers).Error; err != nil {
			return nil, err
		}
		if len(followers) == 0 {
			return &follow.GetFollowerListResponse{}, nil
		}
		var firstPageFollows []*model.Follow
		if in.PageSize < 0 {
			in.PageSize = int64(len(followers))
		}
		if len(followers) > int(in.PageSize) {
			firstPageFollows = followers[:in.PageSize]
		} else {
			firstPageFollows = followers
			isEnd = true
		}
		for _, f := range firstPageFollows {
			followersIdList = append(followersIdList, f.FollowerID)
			curPage = append(curPage, &follow.FollowItem{
				Id:         f.Id,
				UserId:     f.FollowerID,
				CreateTime: f.CreatedAt.Unix(),
			})
		}
	}
	if len(curPage) > 0 {
		last := curPage[len(curPage)-1]
		lastId = last.Id
		cursor = last.CreateTime
		if cursor < 0 {
			cursor = 0
		}
		for k, v := range curPage {
			if in.Cursor == v.CreateTime && v.Id == in.Id {
				curPage = curPage[k:]
				break
			}
		}
	}

	if !isCache {
		if len(followers) < 10*int(in.PageSize) && len(followers) > 0 {
			followers = append(followers, &model.Follow{
				FollowerID: -1,
			})
		}
		// 缓存中没有数据，将数据存入缓存
		for _, v := range followers {
			l.svcCtx.RDB.ZAdd(l.ctx, key, redis.Z{
				Score:  float64(v.CreatedAt.Unix()),
				Member: v.FollowerID,
			})
		}
	}

	return &follow.GetFollowerListResponse{
		Items:  curPage,
		Cursor: cursor,
		IsEnd:  isEnd,
		LastId: lastId,
	}, nil
}

func (l *ListFollowersLogic) cacheFollowers(followKey string, cursor, ps int64) ([]int64, error) {
	maxScore := "inf"
	if cursor > 0 {
		maxScore = fmt.Sprintf("%d", cursor)
	}
	zs, err := l.svcCtx.RDB.ZRevRangeByScore(l.ctx, followKey, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    maxScore,
		Count:  ps,
		Offset: 0,
	}).Result()
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(zs))
	for _, v := range zs {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// 根据id批量查询粉丝信息
func (l *ListFollowersLogic) getFollowersListByIds(userId int64, followerIds []int64) ([]*model.Follow, error) {
	key := fmt.Sprintf("follow_model:%d", userId)
	mkeys := make([]string, 0, len(followerIds))
	for _, v := range followerIds {
		mkeys = append(mkeys, fmt.Sprintf("%d", v))
	}
	result := l.svcCtx.RDB.HMGet(l.ctx, key, mkeys...).Val()
	followSet := make(map[int64]*model.Follow)
	for _, v := range result {
		if v == nil {
			continue
		}
		if v == "" {
			continue
		}
		follow := &model.Follow{}
		err := json.Unmarshal([]byte(v.(string)), follow)
		if err != nil {
			return nil, err
		}
		followSet[follow.FollowerID] = follow
	}
	noHit := make([]int64, 0)
	for _, v := range followerIds {
		if _, ok := followSet[v]; !ok {
			noHit = append(noHit, v)
		}
	}
	// 如果缓存中的获取的数据不够
	if len(noHit) > 0 {
		var list []*model.Follow
		// 在数据库中查询缓存没有的数据，并且存入缓存
		err := l.svcCtx.DB.Model(&model.Follow{}).
			Where("followee_id = ? and follower_id in ?", userId, noHit).
			Find(&list).Error
		if err != nil {
			return nil, err
		}
		for _, v := range list {
			followSet[v.FollowerID] = v
			var b []byte
			if b, err = json.Marshal(v); err != nil {
				return nil, err
			}
			_ = l.svcCtx.RDB.HSet(l.ctx, key, fmt.Sprintf("%d", v.FollowerID), string(b)).Err()
		}
	}
	// 最后组装
	var resultList []*model.Follow
	for _, v := range followerIds {
		if _follow, ok := followSet[v]; ok {
			resultList = append(resultList, _follow)
		}
	}
	return resultList, nil
}
