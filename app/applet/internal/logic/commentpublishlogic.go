package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/comment/commentclient"

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

func (l *CommentPublishLogic) CommentPublish(req *types.CommentPublishRequest, userId int64) (resp *types.CommentPublishResponse, err error) {
	publishCommentResponse, err := l.svcCtx.CommentRPC.PublishComment(l.ctx, &commentclient.PublishCommentRequest{
		BizId:          "video",
		ObjId:          req.VideoId,
		Content:        req.Content,
		ReplayUserId:   userId,
		BeReplayUserId: req.BeReplayUserId,
		ParentId:       req.SuperCommentId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.CommentPublishResponse{
		CommentId: publishCommentResponse.Id,
	}
	return
}
