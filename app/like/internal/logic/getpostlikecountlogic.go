package logic

import (
	"context"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLikeCountLogic {
	return &GetPostLikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某个稿件的点赞数
func (l *GetPostLikeCountLogic) GetPostLikeCount(in *like.GetPostLikeCountRequest) (*like.GetPostLikeCountResponse, error) {
	// todo: add your logic here and delete this line

	return &like.GetPostLikeCountResponse{}, nil
}
