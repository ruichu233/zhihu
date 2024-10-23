package logic

import (
	"context"
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
	var videoModel model.Videos
	if err := l.svcCtx.DB.Model(&model.Videos{}).Limit(1).First(&videoModel, in.VideoId).Error; err != nil {
		return nil, err
	}
	return &video.DetailResponse{
		AuthorId:     videoModel.AuthorId,
		CommentCount: videoModel.CommentNum,
		CoverUrl:     videoModel.CoverUrl,
		Description:  videoModel.Description,
		LikeCount:    videoModel.LikeNum,
		TagIds:       videoModel.TagIds,
		Title:        videoModel.Title,
		VideoUrl:     videoModel.VideoUrl,
	}, nil
}
