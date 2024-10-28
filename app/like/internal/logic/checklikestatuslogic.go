package logic

import (
	"context"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLikeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLikeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLikeStatusLogic {
	return &CheckLikeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询是否对单obj点过赞
func (l *CheckLikeStatusLogic) CheckLikeStatus(in *like.CheckLikeStatusRequest) (*like.CheckLikeStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &like.CheckLikeStatusResponse{}, nil
}
