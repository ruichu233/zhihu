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

// CheckLikeStatus 查询是否对单obj点过赞
func (l *CheckLikeStatusLogic) CheckLikeStatus(in *like.CheckLikeStatusRequest) (*like.CheckLikeStatusResponse, error) {
	_, err := IsLike(l.ctx, l.svcCtx.RDB, l.svcCtx.DB, in.BizId, in.UserId, in.ObjId)
	if err != nil {
		return nil, err
	}
	return &like.CheckLikeStatusResponse{
		BizId:  in.BizId,
		UserId: in.UserId,
		ObjId:  in.ObjId,
	}, nil
}
