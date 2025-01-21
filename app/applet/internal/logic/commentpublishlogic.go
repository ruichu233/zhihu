package logic

import (
	"context"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPublishLogic {
	return &CommentPublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPublishLogic) CommentPublish(req *types.CommentPublishRequest) (resp *types.CommentPublishResponse, err error) {

	return
}
