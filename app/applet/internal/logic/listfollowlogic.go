package logic

import (
	"context"
	"zhihu/app/follow/pb/follow"
	"zhihu/app/user/pb/user"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowLogic {
	return &ListFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowLogic) ListFollow(req *types.ListFollowRequest) (resp *types.ListFollowResponse, err error) {
	resp = &types.ListFollowResponse{
		FollowList: make([]types.FollowInfo, 0),
		IsEnd:      false,
		LastId:     0,
		Cursor:     0,
	}
	var list []*follow.FollowItem
	switch req.FollowType {
	case 1: //朋友
		friendListResponse, err := l.svcCtx.FollowRPC.ListFriends(l.ctx, &follow.GetFriendListRequest{
			UserId:   req.UserId,
			Cursor:   req.Cursor,
			PageSize: req.PageSize,
			Id:       req.LastId,
		})
		if err != nil {
			return nil, err
		}
		list = friendListResponse.Items
	case 2: //关注
		followListResponse, err := l.svcCtx.FollowRPC.ListFollowing(l.ctx, &follow.GetFollowListRequest{
			UserId:   req.UserId,
			Cursor:   req.Cursor,
			PageSize: req.PageSize,
			Id:       0,
		})
		if err != nil {
			return nil, err
		}
		list = followListResponse.Items
	case 3: //粉丝
		followerListResponse, err := l.svcCtx.FollowRPC.ListFollowers(l.ctx, &follow.GetFollowerListRequest{
			UserId:   req.UserId,
			Cursor:   req.Cursor,
			PageSize: req.PageSize,
			Id:       0,
		})
		if err != nil {
			return nil, err
		}
		list = followerListResponse.Items
	}
	// 获取用户id
	userIdList := make([]int64, 0, len(list))
	for _, item := range list {
		userIdList = append(userIdList, item.UserId)
	}
	userList, err := l.svcCtx.UserRPC.GetUserInfoList(l.ctx, &user.UserInfoListRequest{
		UserIdList: userIdList,
	})
	if err != nil {
		return nil, err
	}
	// 获取用户信息集合
	userInfoMap := make(map[int64]*user.UserInfoResponse)
	for _, userInfo := range userList.UserList {
		userInfoMap[userInfo.Id] = userInfo
	}
	// 构建结果
	for _, item := range list {
		userInfo := userInfoMap[item.UserId]
		resp.FollowList = append(resp.FollowList, types.FollowInfo{
			UserId:   item.UserId,
			UserName: userInfo.Username,
			Avatar:   userInfo.Avatar,
		})
	}
	return
}
