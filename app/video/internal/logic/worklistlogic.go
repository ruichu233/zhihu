package logic

import (
	"context"
	"zhihu/app/video/internal/model"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkListLogic {
	return &WorkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据userId获取作品列表
func (l *WorkListLogic) WorkList(in *video.WorkListRequest) (*video.WorkListResponse, error) {
	var videos []*model.Video
	if err := l.svcCtx.DB.Model(&model.Video{}).Find(&videos, "author_id = ?", in.UserId).Error; err != nil {
		return nil, err
	}
	videoFeeds := make([]*video.VideoFeed, 0, len(videos))
	for _, v := range videos {
		videoFeeds = append(videoFeeds, &video.VideoFeed{
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
	return &video.WorkListResponse{
		VideoFeeds: videoFeeds,
	}, nil
}
