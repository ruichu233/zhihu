package logic

import (
	"context"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentDeleteLogic {
	return &CommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentDeleteLogic) CommentDelete(req *types.CommentDeleteRequest) (resp *types.CommentDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
