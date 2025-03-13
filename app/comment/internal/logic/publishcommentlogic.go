package logic

import (
	"context"
	"encoding/json"
	"zhihu/app/comment/model"
	"zhihu/pkg/mq"

	"github.com/yitter/idgenerator-go/idgen"

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
	go func() {
		type notify struct {
			ToUserId   int64  `json:"to_user_id"`
			FromUserId int64  `json:"from_user_id"`
			Type       int32  `json:"type"`
			Content    string `json:"content"`
		}
		n := &notify{
			ToUserId:   in.BeReplayUserId,
			FromUserId: in.ReplayUserId,
			Type:       3, // 评论
			Content:    in.Content,
		}
		value, err := json.Marshal(n)
		if err != nil {
			logx.Error(err)
		}
		if err := l.svcCtx.Producer.Publish("notify_topic", &mq.MsgEntity{
			MsgID: "",
			Key:   "",
			Val:   string(value),
		}); err != nil {
			logx.Error(err)
		}
	}()
	return &comment.PublishCommentResponse{
		Id: _comment.Id,
	}, nil
}
