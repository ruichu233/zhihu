package logic

import (
	"context"
	"zhihu/app/comment/model"

	"zhihu/app/comment/internal/svc"
	"zhihu/app/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取评论列表
func (l *GetCommentListLogic) GetCommentList(in *comment.GetCommentListRequest) (*comment.GetCommentListResponse, error) {
	// 1、获取评论列表
	var comments []*model.Comments
	if err := l.svcCtx.DB.Where("biz_id = ? and obj_id = ? and status = 1", in.BizId, in.ObjId).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	// 2、获取评论点赞数
	likeNumMap := make(map[int64]int64)
	for _, _comment := range comments {
		_comment.LikeNum = likeNumMap[_comment.Id]
	}
	// 3、构建评论层级结构
	var rootComments []*comment.CommentInfo                 // 存放根评论
	var commentMap = make(map[int64][]*comment.CommentInfo) // 存放子评论，key为父评论id，value为子评论
	// 3.1 将评论按照ParentId进行分类
	for _, _comment := range comments {
		// 根评论
		if _comment.ParentId == 0 {
			rootComments = append(rootComments, &comment.CommentInfo{
				Id:             _comment.Id,
				Content:        _comment.Content,
				LikeNum:        _comment.LikeNum,
				ReplayUserId:   _comment.ReplyUserId,
				BeReplayUserId: _comment.BeReplyUserId,
				CreateTime:     _comment.CreatedAt,
				ObjId:          _comment.ObjId,
				FatherId:       _comment.ParentId,
				ReplayList:     nil,
			})
		} else {
			commentMap[_comment.ParentId] = append(commentMap[_comment.ParentId], &comment.CommentInfo{
				Id:             _comment.Id,
				Content:        _comment.Content,
				LikeNum:        _comment.LikeNum,
				ReplayUserId:   _comment.ReplyUserId,
				BeReplayUserId: _comment.BeReplyUserId,
				CreateTime:     _comment.CreatedAt,
				ObjId:          _comment.ObjId,
				FatherId:       _comment.ParentId,
				ReplayList:     nil,
			})
		}
	}
	// 3.2 构建子评论层级结构
	// 遍历根评论，递归添加子评论
	var buildCommentTree func(comment *comment.CommentInfo)
	buildCommentTree = func(comments *comment.CommentInfo) {
		if children, exists := commentMap[comments.Id]; exists {
			comments.ReplayList = children
			for _, child := range children {
				buildCommentTree(child)
			}
		}
	}
	for _, root := range rootComments {
		buildCommentTree(root)
	}
	// 3.3 构建评论树
	commentTree := make([]*comment.CommentInfo, 0)
	commentTree = append(commentTree, rootComments...)
	return &comment.GetCommentListResponse{
		CommentList: commentTree,
	}, nil
}
