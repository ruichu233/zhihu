package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"zhihu/app/video/internal/model"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailListLogic {
	return &DetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据IdList获取视频详情列表
func (l *DetailListLogic) DetailList(in *video.DetailListRequest) (*video.DetailListResponse, error) {
	var (
		videoSet    = make(map[int64]*model.Video)
		noHitIdList []int64
	)

	pip := l.svcCtx.RDB.Pipeline()
	for _, id := range in.VideoIds {
		pip.Get(l.ctx, fmt.Sprintf("%s%d", "video_", id))
	}
	result, _ := pip.Exec(l.ctx)
	for _, v := range result {
		res := v.(*redis.StringCmd).Val()
		if res != "" {
			_video := &model.Video{}
			err := json.Unmarshal([]byte(res), _video)
			if err != nil {
				return nil, err
			}
			videoSet[_video.Id] = _video
		}
	}
	// 寻找未命中缓存的id
	for _, id := range in.VideoIds {
		if _, ok := videoSet[id]; !ok {
			noHitIdList = append(noHitIdList, id)
		}
	}

	videos := make([]*model.Video, 0, len(noHitIdList))
	if err := l.svcCtx.DB.Model(&model.Video{}).Where("id in ?", videoSet).Find(&videos).Error; err != nil {
		return nil, err
	}
	videoItems := make([]*video.VideoFeed, 0, len(videos))
	for _, v := range videos {
		videoItems = append(videoItems, &video.VideoFeed{
			VideoId:      v.Id,
			AuthorId:     v.AuthorId,
			CommentCount: v.CommentNum,
			CoverUrl:     v.CoverUrl,
			Description:  v.Description,
			LikeCount:    v.LikeNum,
			Title:        v.Title,
			VideoUrl:     v.VideoUrl,
			CreateTime:   v.CreatedAt.Unix(),
		})
	}
	return &video.DetailListResponse{
		VideoFeeds: videoItems,
	}, nil
}
