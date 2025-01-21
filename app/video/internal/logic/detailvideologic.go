package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"zhihu/app/video/internal/model"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailVideoLogic {
	return &DetailVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据Id获取视频详情
func (l *DetailVideoLogic) DetailVideo(in *video.DetailRequest) (*video.DetailResponse, error) {
	// 1、查缓存
	cacheKey := GetVideoKey(in.VideoId)
	result, err := l.svcCtx.RDB.HGetAll(l.ctx, cacheKey).Result()
	if err == nil {
		// 处理结果
		var (
			authorId     int64
			commentCount int64
			likeCount    int64
		)
		for k, v := range result {
			switch k {
			case "author_id":
				authorId, _ = strconv.ParseInt(v, 10, 64)
			case "comment_count":
				commentCount, _ = strconv.ParseInt(v, 10, 64)
			case "like_count":
				likeCount, _ = strconv.ParseInt(v, 10, 64)
			}

			return &video.DetailResponse{
				AuthorId:     authorId,
				CommentCount: commentCount,
				CoverUrl:     result["cover_url"],
				Description:  result["description"],
				LikeCount:    likeCount,
				Title:        result["title"],
				VideoUrl:     result["video_url"],
			}, nil
		}

		var videoModel model.Video
		if err := l.svcCtx.DB.Model(&model.Video{}).Limit(1).First(&videoModel, in.VideoId).Error; err != nil {
			return nil, err
		}
		// 更新缓存
		go func() {
			mp := make(map[string]interface{})
			mp["author_id"] = videoModel.AuthorId
			mp["comment_count"] = videoModel.CommentNum
			mp["like_count"] = videoModel.LikeNum
			mp["cover_url"] = videoModel.CoverUrl
			mp["description"] = videoModel.Description
			mp["title"] = videoModel.Title
			mp["video_url"] = videoModel.VideoUrl
			l.svcCtx.RDB.HSet(l.ctx, cacheKey, mp)
			l.svcCtx.RDB.Expire(l.ctx, cacheKey, 60*time.Minute)
		}()

		data := &video.DetailResponse{
			AuthorId:     videoModel.AuthorId,
			CommentCount: videoModel.CommentNum,
			CoverUrl:     videoModel.CoverUrl,
			Description:  videoModel.Description,
			LikeCount:    videoModel.LikeNum,
			Title:        videoModel.Title,
			VideoUrl:     videoModel.VideoUrl,
		}
		return data, nil
	}
	return nil, err
}

func GetVideoKey(videoId int64) string {
	return fmt.Sprintf("VIDEO_%d", videoId)
}
