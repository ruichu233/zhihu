package logic

import (
	"context"
	"fmt"
	"time"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUploadURLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUploadURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUploadURLLogic {
	return &GetUploadURLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUploadURL 获取视频上传的预签名 URL
func (l *GetUploadURLLogic) GetUploadURL(in *video.GetUploadURLRequest) (*video.GetUploadURLResponse, error) {
	objectName := fmt.Sprintf("%d/%s", time.Now().Unix(), in.Filename)
	videoUrl, err := l.svcCtx.OSS.PresignedPutObject(l.ctx, "video", objectName+".mp4", 5*time.Minute)
	if err != nil {
		return nil, err
	}
	coverUrl, err := l.svcCtx.OSS.PresignedPutObject(l.ctx, "cover", objectName+".jpg", 5*time.Minute)
	if err != nil {
		return nil, err
	}
	return &video.GetUploadURLResponse{
		VideoUrl: videoUrl.String(),
		CoverUrl: coverUrl.String(),
	}, nil
}
