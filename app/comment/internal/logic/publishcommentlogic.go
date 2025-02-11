package logic

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"zhihu/app/comment/model"

	"zhihu/app/comment/internal/svc"
	"zhihu/app/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishCommentLogic {
	return &PublishCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布评论
func (l *PublishCommentLogic) PublishComment(in *comment.PublishCommentRequest) (*comment.PublishCommentResponse, error) {
	// 判断是否为回复评论
	if in.ParentId != 0 {
		var parentComment model.Comments
		if err := l.svcCtx.DB.Where("id=?", in.ParentId).First(&parentComment).Error; err != nil {
			return nil, err
		}
	}
	// 存储评论
	_comment := &model.Comments{
		BaseModel: model.BaseModel{
			Id: idgen.NextId(),
		},
		BizId:         in.BizId,
		ObjId:         in.ObjId,
		ReplyUserId:   in.ReplayUserId,
		BeReplyUserId: in.BeReplayUserId,
		ParentId:      in.ParentId,
		Content:       in.Content,
		Status:        1,
		LikeNum:       0,
	}
	if err := l.svcCtx.DB.Create(_comment).Error; err != nil {
		return nil, err
	}
	// 异步通知（如果是回复则通知被回复人）
	if in.ParentId != 0 {
		// 通知被回复人
		// todo
	}
	return &comment.PublishCommentResponse{
		Id: _comment.Id,
	}, nil
}
