package logic

import (
	"context"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLikersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLikersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLikersLogic {
	return &GetPostLikersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询点赞人列表
func (l *GetPostLikersLogic) GetPostLikers(in *like.GetPostLikersRequest) (*like.GetPostLikersResponse, error) {
	// todo: add your logic here and delete this line

	return &like.GetPostLikersResponse{}, nil
}
