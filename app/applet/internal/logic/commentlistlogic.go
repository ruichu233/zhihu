package logic

import (
	"context"
	"zhihu/app/comment/commentclient"
	"zhihu/app/comment/pb/comment"
	"zhihu/app/user/userclient"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	// 1、从评论服务获取评论树
	commentListResp, err := l.svcCtx.CommentRPC.GetCommentList(l.ctx, &commentclient.GetCommentListRequest{
		BizId: "video",
		ObjId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	// 2、递归评论树获取 userId 列表
	userIdSet := make(map[int64]struct{})
	for _, commentInfo := range commentListResp.CommentList {
		getUserIdFromCommentItem(commentInfo, userIdSet)
	}
	userIdList := make([]int64, 0, len(userIdSet))
	for userId := range userIdSet {
		userIdList = append(userIdList, userId)
	}
	// 3、从用户服务获取用户信息
	userInfoListResp, err := l.svcCtx.UserRPC.GetUserInfoList(l.ctx, &userclient.UserInfoListRequest{
		UserIdList: userIdList,
	})
	if err != nil {
		return nil, err
	}
	authorSet := make(map[int64]*userclient.UserInfoResponse)
	for _, userInfo := range userInfoListResp.UserList {
		authorSet[userInfo.Id] = userInfo
	}
	// 4、组装返回结果
	commentList := make([]types.CommentInfo, 0, len(commentListResp.CommentList))
	for _, commentInfo := range commentListResp.CommentList {
		comment := &types.CommentInfo{}
		buildCommentInfo(commentInfo, comment, authorSet)
		commentList = append(commentList, *comment)
	}
	resp = &types.CommentListResponse{
		CommentList: commentList,
	}
	return
}

func getUserIdFromCommentItem(c *comment.CommentInfo, userIdSet map[int64]struct{}) {
	for _, commentInfo := range c.ReplayList {
		getUserIdFromCommentItem(commentInfo, userIdSet)
	}
	userIdSet[c.ReplayUserId] = struct{}{}
	userIdSet[c.BeReplayUserId] = struct{}{}
}

func buildCommentInfo(c *comment.CommentInfo, comment *types.CommentInfo, userMap map[int64]*userclient.UserInfoResponse) *types.CommentInfo {

	for _, commentInfo := range c.ReplayList {
		comment.Children = append(comment.Children, *buildCommentInfo(commentInfo, comment, userMap))
	}
	if len(comment.Children) == 0 {
		comment.Children = make([]types.CommentInfo, 0)
	}
	umap := userMap
	userInfo := umap[c.ReplayUserId]
	if userInfo == nil {
		userInfo = &userclient.UserInfoResponse{}
	}
	comment.ObjId = c.ObjId
	comment.Id = c.Id
	comment.Content = c.Content
	comment.AddTime = c.CreateTime
	comment.LikeNums = c.LikeNum
	comment.Avatar = userInfo.Avatar
	comment.UserId = c.ReplayUserId
	comment.BeReplayUserId = c.BeReplayUserId
	comment.Nickname = userInfo.Username
	comment.SuperNickname = userMap[c.BeReplayUserId].Username
	comment.SuperCommentId = c.FatherId
	comment.IsLike = 0
	return comment
}
