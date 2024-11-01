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
	// 构建缓存键
	Key := GetFollowKey(in.UserId)
	// 查询缓存

	return &follow.GetFollowerListResponse{}, nil
}

func (l *ListFollowersLogic) cacheFollowers(followKey string) error {

	return nil
}
