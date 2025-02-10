package logic

import (
	"context"
	"zhihu/app/comment/commentclient"

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
	_, err = l.svcCtx.CommentRPC.DeleteComment(l.ctx, &commentclient.DeleteCommentRequest{
		Id: req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CommentDeleteResponse{
		Status: "Success",
	}, nil
}
