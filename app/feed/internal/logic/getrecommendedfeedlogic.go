package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"zhihu/app/feed/internal/svc"
	"zhihu/app/feed/pb/feed"
)

type GetRecommendedFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecommendedFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendedFeedLogic {
	return &GetRecommendedFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取个性化推荐 Feed
func (l *GetRecommendedFeedLogic) GetRecommendedFeed(in *feed.GetRecommendedFeedRequest) (*feed.GetRecommendedFeedResponse, error) {
	userIdStr := fmt.Sprintf("%d", in.UserId)
	recommend, err := l.svcCtx.Gorse.GetRecommend(l.ctx, userIdStr, "", int(in.PageSize), int(in.Page*in.PageSize))
	if err != nil {
		return nil, err
	}
	// 已读反馈
	//feedbackList := make([]client.Feedback, 0)
	//for _, v := range recommend {
	//	feedbackList = append(feedbackList, client.Feedback{
	//		FeedbackType: "read",
	//		ItemId:       v,
	//		UserId:       userIdStr,
	//		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
	//	})
	//}
	//_, err = l.svcCtx.Gorse.InsertFeedback(l.ctx, feedbackList)
	//if err != nil {
	//	return nil, err
	//}
	items := make([]int64, 0)

	for _, v := range recommend {
		item, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return &feed.GetRecommendedFeedResponse{
		RecommendedItems: items,
	}, nil
}
