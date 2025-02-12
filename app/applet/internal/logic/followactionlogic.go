package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowActionLogic) FollowAction(req *types.FollowActionRequest, userId int64) (resp *types.FollowActionResponse, err error) {
	resp = &types.FollowActionResponse{
		Status: "success",
	}
	followActionResponse, err := l.svcCtx.FollowRPC.FollowAction(l.ctx, &follow.FollowActionRequest{
		ActionType: follow.FollowActionRequest_ActionType(req.ActionType),
		FolloweeId: req.FolloweeId,
		FollowerId: userId,
	})
	if err != nil {
		return nil, err
	}
	if !followActionResponse.Success {
		resp.Status = "failed"
	}
	return
}
