package logic

import (
	"context"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeListLogic {
	return &LikeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据userId获取喜欢列表
func (l *LikeListLogic) LikeList(in *video.LikeListRequest) (*video.LikeListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.LikeListResponse{}, nil
}
