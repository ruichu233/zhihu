package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"zhihu/app/follow/model"

	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/pb/follow"

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

	return &follow.GetFriendListResponse{}, nil
}

// 根据id批量查询关注信息
func (l *ListFollowingLogic) getFriendListByIds(userId int64, friendIds []int64) ([]*model.Follow, error) {
	key := fmt.Sprintf("friend:%d", userId)
	mkeys := make([]string, 0, len(friendIds))
	for _, v := range friendIds {
		mkeys = append(mkeys, fmt.Sprintf("%d", v))
	}
	result, _ := l.svcCtx.RDB.HMGet(l.ctx, key, mkeys...).Result()
	followSet := make(map[int64]*model.Follow)
	for _, v := range result {
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
	for _, v := range friendIds {
		if _follow, ok := followSet[v]; ok {
			resultList = append(resultList, _follow)
		}
	}
	return resultList, nil
}
