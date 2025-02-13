package logic

import (
	"context"
	"zhihu/app/user/pb/user"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAvatarLogic {
	return &GetAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAvatarLogic) GetAvatar(req *types.GetAvatarRequest) (resp *types.GetAvatarResponse, err error) {
	avatarResponse, err := l.svcCtx.UserRPC.GetAVatarUrl(l.ctx, &user.GetAvatarRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return
	}
	resp = &types.GetAvatarResponse{
		AvatarUrl: avatarResponse.AvatarUrl,
	}
	return
}
