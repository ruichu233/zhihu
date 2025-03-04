package logic

import (
	"context"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowLogic {
	return &CheckFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFollowLogic) CheckFollow(req *types.CheckFollowRequest) (resp *types.CheckFollowResponse, err error) {
	isFollowResp, err := l.svcCtx.FollowRPC.IsFollow(l.ctx, &follow.IsFollowRequest{
		UserId:   req.UserId,
		ToUserId: req.ToUserId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.CheckFollowResponse{
		IsFollow: isFollowResp.IsFollow,
	}
	return
}
