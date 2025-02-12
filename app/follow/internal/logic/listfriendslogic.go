package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/model"
	"zhihu/app/follow/pb/follow"
	"zhihu/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendsLogic {
	return &ListFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListFriendsLogic) ListFriends(in *follow.GetFriendListRequest) (*follow.GetFriendListResponse, error) {
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
	friendKey := GetFriendKey(in.UserId)
	// 2、从缓存中获取关注id列表
	followIdList, _ = l.cacheFriend(friendKey, in.Cursor, in.PageSize)
	if len(followIdList) > 0 {
		isCache = true
		if followIdList[len(followIdList)-1] == -1 {
			followIdList = followIdList[:len(followIdList)-1]
			isEnd = true
		}
		if len(followIdList) == 0 {
			return &follow.GetFriendListResponse{}, nil
		}
		_follows, err := l.getFriendListByIds(in.UserId, followIdList)
		if err != nil {
			return nil, err
		}
		follows = _follows
		for _, v := range follows {
			curPage = append(curPage, &follow.FollowItem{
				Id:         v.Id,
				UserId:     v.FollowerID,
				CreateTime: v.CreatedAt.Unix(),
			})
		}
	} else {
		defaultCache := 10 * in.PageSize
		friendsId, err := getFriends(l.svcCtx.DB, in.UserId, utils.ParseTimestamp(in.Cursor), defaultCache)
		if err != nil {
			return nil, err
		}
		if len(friendsId) == 0 {
			return &follow.GetFriendListResponse{}, nil
		}
		if err = l.svcCtx.DB.Model(&model.Follow{}).Where("followee_id = ? and follower_id in ?", in.UserId, friendsId).First(&follows).Error; err != nil {
			return nil, err
		}
		var firstPageFollows []*model.Follow
		if len(follows) > int(in.PageSize) {
			firstPageFollows = follows[:in.PageSize]
		} else {
			firstPageFollows = follows
			isEnd = true
		}
		for _, f := range firstPageFollows {
			followIdList = append(followIdList, f.FollowerID)
			curPage = append(curPage, &follow.FollowItem{
				Id:         f.Id,
				UserId:     f.FollowerID,
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
					FollowerID: -1,
				})
				for _, v := range follows {
					l.svcCtx.RDB.ZAdd(l.ctx, friendKey, redis.Z{
						Score:  float64(v.UpdatedAt.Unix()),
						Member: v.FollowerID,
					})
				}
			}
		}()
	}

	return &follow.GetFriendListResponse{
		Items:  curPage,
		Cursor: cursor,
		IsEnd:  isEnd,
		LastId: lastId,
	}, nil
}
func (l *ListFriendsLogic) cacheFriend(key string, cursor, ps int64) ([]int64, error) {
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
func (l *ListFriendsLogic) getFriendListByIds(userId int64, friendIds []int64) ([]*model.Follow, error) {
	key := fmt.Sprintf("follow_model:%d", userId)
	mkeys := make([]string, 0, len(friendIds))
	for _, v := range friendIds {
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
		followSet[_follow.FollowerID] = _follow
	}
	noHit := make([]int64, 0)
	for _, v := range friendIds {
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
	for _, v := range friendIds {
		if _follow, ok := followSet[v]; ok {
			resultList = append(resultList, _follow)
		}
	}
	return resultList, nil
}

func getFriends(db *gorm.DB, userA int64, cretedAt string, pageSize int64) ([]int64, error) {
	var friendIDs []int64
	// 执行 SQL 查询，获取 A 的所有朋友
	sql := `
		SELECT f1.followee_id AS friend_id
		FROM follows f1
		JOIN follows f2 ON f1.followee_id = f2.follower_id
		WHERE f1.follower_id = ? AND f2.followee_id = ? AND f1.deleted_at IS NULL AND f1.created_at < ?
	`
	var err error
	if pageSize > 0 {
		sql += " LIMIT ?"
		err = db.Raw(sql, userA, userA, cretedAt, pageSize).Pluck("friend_id", &friendIDs).Error
	} else {
		err = db.Raw(sql, userA, userA, cretedAt).Pluck("friend_id", &friendIDs).Error
	}
	if err != nil {
		return nil, err
	}
	return friendIDs, nil
}
