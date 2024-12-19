package logic

import (
	"context"

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

func (l *ListFollowersLogic) ListFollowers(in *follow.GetFollowerListRequest) (*follow.GetFollowerListResponse, error) {
	if in.PageSize < 0 {
		in.PageSize = 10
	}
	if in.Cursor < 0 {
		in.Cursor = 0
	}
	var (
		curPage []*follow.FollowerItem
	)

	// 构建缓存键
	Key := GetFollowerKey(in.UserId)
	// 查询缓存
	err, followers := l.cacheFollowers(Key, in.Cursor, in.PageSize)
	return &follow.GetFollowerListResponse{}, nil
}

func (l *ListFollowersLogic) cacheFollowers(followKey string, cursor, ps int64) (error, []int64) {

	return nil
}
