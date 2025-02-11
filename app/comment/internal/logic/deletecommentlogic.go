package logic

import (
	"context"
	"zhihu/app/comment/internal/svc"
	"zhihu/app/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除评论
func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	if err := l.svcCtx.DB.Where("id = ?", in.Id).Update("status", 0).Error; err != nil {
		return nil, err
	}
	return &comment.DeleteCommentResponse{
		Success: true,
	}, nil
}
