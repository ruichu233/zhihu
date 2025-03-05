package logic

import (
	"context"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/feed/pb/feed"
	"zhihu/app/follow/pb/follow"
	"zhihu/app/like/pb/like"
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
			PageSize: req.PageSize,
		})
		if err != nil {
			return nil, err
		}
		feedList = resp.Items
	}
	if req.FeedType == 2 {
		resp, err := l.svcCtx.FeedRPC.GetRecommendedFeed(l.ctx, &feed.GetRecommendedFeedRequest{
			UserId:   0,
			Page:     0,
			PageSize: 10,
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
		})
		if err != nil {
			return nil, err
		}
		// 获取用户的关注列表
		followListResp, err := l.svcCtx.FollowRPC.ListFollowing(l.ctx, &follow.GetFollowListRequest{
			UserId:   userId,
			Cursor:   0,
			PageSize: -1,
			Id:       0,
		})
		if err != nil {
			return nil, err
		}
		followMap := make(map[int64]bool)
		for _, follow := range followListResp.Items {
			followMap[follow.UserId] = true
		}

		for _, videoFeed := range detailListResp.VideoFeeds {
			user, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &userclient.UserInfoRequest{
				UserId: videoFeed.AuthorId,
			})
			if err != nil {
				return nil, err
			}
			status, err := l.svcCtx.LikeRPC.CheckLikeStatus(l.ctx, &like.CheckLikeStatusRequest{
				BizId:  "video_like",
				ObjId:  videoFeed.VideoId,
				UserId: userId,
			})
			if err != nil {
				return nil, err
			}

			countResp, err := l.svcCtx.LikeRPC.GetPostLikeCount(l.ctx, &like.GetPostLikeCountRequest{
				BizId: "video_like",
				ObjId: videoFeed.VideoId,
			})
			if err != nil {
				return nil, err
			}

			videoInfo := types.VideoInfo{
				VideoId:       videoFeed.VideoId,
				AuthorId:      videoFeed.AuthorId,
				AuthorName:    user.Username,
				AuthorAvatar:  user.Avatar,
				Title:         videoFeed.Title,
				VideoUrl:      videoFeed.VideoUrl,
				CoverUrl:      videoFeed.CoverUrl,
				Description:   videoFeed.Description,
				CommentCount:  videoFeed.CommentCount,
				LikeCount:     countResp.Count,
				IsLike:        status.IsLiked,
				IsInteraction: followMap[videoFeed.AuthorId],
			}
			resp.VideoList = append(resp.VideoList, videoInfo)
		}
	}
	return resp, nil
}
