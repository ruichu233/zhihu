package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"zhihu/app/video/internal/model"
	"zhihu/app/video/pb/video"

	"zhihu/app/video/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *video.DetailRequest) (*video.DetailResponse, error) {
	// 1、查缓存
	cacheKey := GetVideoKey(in.VideoId)
	result, err := l.svcCtx.RDB.HGetAll(l.ctx, cacheKey).Result()
	if err == nil {
		// 处理结果
		var (
			authorId     int64
			commentCount int64
			likeCount    int64
			tagIds       []int64
		)
		for k, v := range result {
			switch k {
			case "author_id":
				authorId, _ = strconv.ParseInt(v, 10, 64)
			case "comment_count":
				commentCount, _ = strconv.ParseInt(v, 10, 64)
			case "like_count":
				likeCount, _ = strconv.ParseInt(v, 10, 64)
			case "tag_ids":
				split := strings.Split(v, "|")
				for _, v := range split {
					id, _ := strconv.ParseInt(v, 10, 64)
					tagIds = append(tagIds, id)
				}
			}
		}

		return &video.DetailResponse{
			AuthorId:     authorId,
			CommentCount: commentCount,
			CoverUrl:     result["cover_url"],
			Description:  result["description"],
			LikeCount:    likeCount,
			TagIds:       tagIds,
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
		tagIdStrs := make([]string, 0, len(videoModel.TagIds))
		for _, v := range videoModel.TagIds {
			tagIdStrs = append(tagIdStrs, strconv.FormatInt(v, 10))
		}
		mp["author_id"] = videoModel.AuthorId
		mp["comment_count"] = videoModel.CommentNum
		mp["like_count"] = videoModel.LikeNum
		mp["tag_ids"] = strings.Join(tagIdStrs, "|")
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
		TagIds:       videoModel.TagIds,
		Title:        videoModel.Title,
		VideoUrl:     videoModel.VideoUrl,
	}
	return data, nil
}

func GetVideoKey(videoId int64) string {
	return fmt.Sprintf("VIDEO_%d", videoId)
}
