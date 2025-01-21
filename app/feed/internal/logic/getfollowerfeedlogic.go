package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"zhihu/app/follow/pb/follow"
	"zhihu/app/video/pb/video"

	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedItem struct {
	Member int64
	Score  float64
}

type GetFollowerFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerFeedLogic {
	return &GetFollowerFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注者的 Feed
func (l *GetFollowerFeedLogic) GetFollowerFeed(in *feed.GetFollowerFeedRequest) (*feed.GetFollowerFeedResponse, error) {
	if in.UserId == 0 {
		return nil, errors.New("user_id不能为空")
	}
	if in.PageSize == 0 {
		in.PageSize = 10
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}

	var (
		err     error
		isEnd   bool
		cursor  int64
		curPage []int64
	)
	// 1、缓存
	followerFeed, err := l.cacheFollowerFeed(l.ctx, in.UserId, in.Cursor, in.PageSize)
	if err != nil {
		return nil, err
	}
	if len(followerFeed) > 0 {
		lastItem := followerFeed[len(followerFeed)-1]
		cursor = int64(lastItem.Score)
		if len(followerFeed) < int(in.PageSize) {
			isEnd = true
		}
		for _, item := range followerFeed {
			curPage = append(curPage, item.Member)
		}
	}
	return &feed.GetFollowerFeedResponse{
		Items:      curPage,
		NextCursor: cursor,
		HasMore:    !isEnd,
	}, nil
}

func (l *GetFollowerFeedLogic) cacheFollowerFeed(ctx context.Context, userId int64, cursor int64, pageSize int64) ([]FeedItem, error) {
	key := GetCacheKey(userId)
	b, err := l.svcCtx.RDB.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if b > 0 {
		if err := l.svcCtx.RDB.Expire(ctx, key, time.Hour*24).Err(); err != nil {
			return nil, err
		}
	} else {
		// 不存在说明已经过期，重新构建缓存
		// 1、遍历所有关注的人
		followListResponse, err := l.svcCtx.FollowRPC.ListFollowing(ctx, &follow.GetFollowListRequest{
			UserId:   userId,
			Cursor:   -1,
			PageSize: -1,
			Id:       0,
		})
		if err != nil {
			return nil, err
		}
		// 2、获取作品列表
		for _, following := range followListResponse.Items {
			workListResponse, err := l.svcCtx.VideoRPC.WorkList(ctx, &video.WorkListRequest{
				UserId: following.FolloweeId,
			})
			if err != nil {
				return nil, err
			}
			// 存入redis
			for _, videoFeed := range workListResponse.VideoFeeds {
				l.svcCtx.RDB.ZAdd(ctx, GetCacheKey(userId), redis.Z{
					Member: videoFeed.VideoId,
					Score:  float64(videoFeed.CreateTime),
				})
			}
		}
	}
	res, err := l.svcCtx.RDB.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    "0",
		Max:    fmt.Sprintf("%d", cursor),
		Count:  pageSize,
		Offset: 0,
	}).Result()
	if err != nil {
		return nil, err
	}
	feeds := make([]FeedItem, 0, len(res))
	for _, v := range res {
		item := FeedItem{}
		item.Member = v.Member.(int64)
		item.Score = v.Score
		feeds = append(feeds, item)
	}
	return feeds, nil
}
