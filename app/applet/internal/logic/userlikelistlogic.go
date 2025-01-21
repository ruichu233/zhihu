package logic

import (
	"context"
	"zhihu/app/like/pb/like"
	"zhihu/app/user/pb/user"
	"zhihu/app/user/userclient"
	"zhihu/app/video/pb/video"

	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikeListLogic {
	return &UserLikeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLikeListLogic) UserLikeList(req *types.UserLikeListRequest) (resp *types.UserLikeListResponse, err error) {
	resp = &types.UserLikeListResponse{}
	// 1、获取用户点赞视频id列表
	getUserLikesResponse, err := l.svcCtx.LikeRPC.GetUserLikes(l.ctx, &like.GetUserLikesRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	// 2、通过商品id列表获取视频详情
	videoList, err := l.svcCtx.VideoRPC.DetailList(l.ctx, &video.DetailListRequest{VideoIds: getUserLikesResponse.PostIds})
	if err != nil {
		return nil, err
	}
	// 3、获取视频作者id列表并且根据列表获取作者信息
	authorIdList := make([]int64, 0, len(videoList.VideoFeeds))
	authorSet := make(map[int64]*user.UserInfoResponse, len(videoList.VideoFeeds))
	for _, videoFeed := range videoList.VideoFeeds {
		authorIdList = append(authorIdList, videoFeed.AuthorId)
	}
	GetUserInfoListResp, err := l.svcCtx.UserRPC.GetUserInfoList(l.ctx, &userclient.UserInfoListRequest{
		UserIdList: authorIdList,
	})
	if err != nil {
		return nil, err
	}
	for _, userInfo := range GetUserInfoListResp.UserList {
		authorSet[userInfo.Id] = userInfo
	}
	// 4、组装返回结果
	resp.VideoList = make([]types.VideoInfo, 0, len(videoList.VideoFeeds))
	for _, videoFeed := range videoList.VideoFeeds {
		resp.VideoList = append(resp.VideoList, types.VideoInfo{
			VideoId:      videoFeed.VideoId,
			AuthorId:     videoFeed.AuthorId,
			AuthorAvatar: authorSet[videoFeed.AuthorId].Avatar,
			AuthorName:   authorSet[videoFeed.AuthorId].Username,
			Title:        videoFeed.Title,
			VideoUrl:     videoFeed.VideoUrl,
			CoverUrl:     videoFeed.CoverUrl,
			Description:  videoFeed.Description,
			CommentCount: videoFeed.CommentCount,
			LikeCount:    videoFeed.LikeCount,
		})
	}
	return
}
