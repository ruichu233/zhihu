package logic

import (
	"context"
	"zhihu/app/user/model"

	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoListLogic {
	return &GetUserInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoListLogic) GetUserInfoList(in *user.UserInfoListRequest) (*user.UserInfoListResponse, error) {

	var userList []*model.User
	if err := l.svcCtx.DB.Model(&model.User{}).Where("id in ?", in.UserIdList).Find(&userList).Error; err != nil {
		return nil, err
	}
	resp := &user.UserInfoListResponse{}
	for _, u := range userList {
		userInfo := &user.UserInfoResponse{
			Id:            u.Id,
			Username:      u.Username,
			Email:         u.Email,
			Avatar:        u.Avatar,
			Signature:     u.Signature,
			FollowerCount: u.FollowerCount,
			FollowedCount: u.FollowCount,
		}
		resp.UserList = append(resp.UserList, userInfo)
	}

	return resp, nil
}
