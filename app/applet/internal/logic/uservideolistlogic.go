package logic

import (
	"context"
	"zhihu/app/user/userclient"
	"zhihu/app/video/pb/video"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserVideoList 获取用户作品列表
func (l *UserVideoListLogic) UserVideoList(req *types.UserVideoListRequest) (resp *types.UserVideoListResponse, err error) {
	resp = &types.UserVideoListResponse{}
	// 1、获取用户信息
	userInfoResponse, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &userclient.UserInfoRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	// 2、获取用户作品Id列表
	workListResp, err := l.svcCtx.VideoRPC.WorkList(l.ctx, &video.WorkListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	// 3、组合响应结果
	resp.VideoList = make([]types.VideoInfo, 0, len(workListResp.VideoFeeds))
	for _, videoFeed := range workListResp.VideoFeeds {
		resp.VideoList = append(resp.VideoList, types.VideoInfo{
			VideoId:      videoFeed.VideoId,
			AuthorId:     videoFeed.AuthorId,
			AuthorName:   userInfoResponse.Username,
			AuthorAvatar: userInfoResponse.Avatar,
			VideoUrl:     videoFeed.VideoUrl,
			CoverUrl:     videoFeed.CoverUrl,
			LikeCount:    videoFeed.LikeCount,
			CommentCount: videoFeed.CommentCount,
			Description:  videoFeed.Description,
		})
	}
	return
}
