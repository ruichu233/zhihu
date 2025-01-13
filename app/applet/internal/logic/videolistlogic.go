package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/feed/pb/feed"
	"zhihu/app/user/userclient"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListRequest, userId int64) (resp *types.VideoListResponse, err error) {
	resp = &types.VideoListResponse{}
	// 1、获取feed列表
	var feedList []int64
	if req.FeedType == 1 {
		resp, err := l.svcCtx.FeedRPC.GetFollowerFeed(l.ctx, &feed.GetFollowerFeedRequest{
			UserId:   userId,
			Cursor:   req.Cursor,
			PageSize: req.Page,
		})
		if err != nil {
			return nil, err
		}
		feedList = resp.Items
	}
	if req.FeedType == 2 {
		resp, err := l.svcCtx.FeedRPC.GetRecommendedFeed(l.ctx, &feed.GetRecommendedFeedRequest{
			UserId:   userId,
			Cursor:   req.Cursor,
			PageSize: req.Page,
		})
		if err != nil {
			return nil, err
		}
		feedList = resp.RecommendedItems
	}
	// 2、根据feed列表获取视频信息
	if len(feedList) > 0 {
		detailListResp, err := l.svcCtx.VideoRPC.DetailList(l.ctx, &video.DetailListRequest{
			VideoIds: feedList,
			FeedType: video.VideoFeedType(req.FeedType),
		})
		if err != nil {
			return nil, err
		}
		for _, videoFeed := range detailListResp.VideoFeeds {
			user, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &userclient.UserInfoRequest{
				UserId: videoFeed.AuthorId,
			})
			if err != nil {
				return nil, err
			}
			videoInfo := types.VideoInfo{
				VideoId:      videoFeed.VideoId,
				AuthorId:     videoFeed.AuthorId,
				AuthorName:   user.Username,
				AuthorAvatar: user.Avatar,
				Title:        videoFeed.Title,
				VideoUrl:     videoFeed.VideoUrl,
				CoverUrl:     videoFeed.CoverUrl,
				Description:  videoFeed.Description,
				CommentCount: videoFeed.CommentCount,
				LikeCount:    videoFeed.LikeCount,
			}
			resp.VideoList = append(resp.VideoList, videoInfo)
		}
	}
	return resp, nil
}
