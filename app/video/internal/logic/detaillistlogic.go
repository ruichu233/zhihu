package logic

import (
	"context"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailListLogic {
	return &DetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据IdList获取视频详情列表
func (l *DetailListLogic) DetailList(in *video.DetailListRequest) (*video.DetailListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.DetailListResponse{}, nil
}
