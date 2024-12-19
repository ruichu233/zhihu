package logic

import (
	"context"
	"zhihu/app/video/videoclient"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUrlLogic {
	return &UploadUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadUrlLogic) UploadUrl(filename string) (resp *types.UploadUrlResponse, err error) {
	getUploadURL, err := l.svcCtx.VideoRPC.GetUploadURL(l.ctx, &videoclient.GetUploadURLRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, err
	}
	return &types.UploadUrlResponse{
		VideoUrl: getUploadURL.VideoUrl,
		CoverUrl: getUploadURL.CoverUrl,
	}, nil
}
