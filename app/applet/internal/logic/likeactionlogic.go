package logic

import (
	"context"
	"zhihu/app/like/pb/like"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeActionLogic {
	return &LikeActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeActionLogic) LikeAction(req *types.LikeActionRequest) error {
	_, err := l.svcCtx.LikeRPC.LikeAction(l.ctx, &like.LikeActionRequest{
		ActionType: like.LikeActionRequest_ActionType(req.ActionType),
		BizId:      req.BizId,
		ObjId:      req.ObjId,
	})
	if err != nil {
		return err
	}
	return nil
}
