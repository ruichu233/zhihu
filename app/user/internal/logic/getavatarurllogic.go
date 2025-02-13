package logic

import (
	"context"
	"strconv"
	"time"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAVatarUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAVatarUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAVatarUrlLogic {
	return &GetAVatarUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAVatarUrlLogic) GetAvatarUrl(in *user.GetAvatarRequest) (*user.GetAvatarResponse, error) {
	avatarUrl, err := l.svcCtx.OSS.PresignedPutObject(l.ctx, "avatar", strconv.FormatInt(in.UserId, 10)+strconv.FormatInt(time.Now().Unix(), 10)+".jpg", 5*time.Minute)
	if err != nil {
		return nil, err
	}
	return &user.GetAvatarResponse{
		AvatarUrl: avatarUrl.String(),
	}, nil
}
