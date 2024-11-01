package logic

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"zhihu/app/feed/pb/feed"
	"zhihu/app/video/internal/model"
	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *video.PublishRequest) (*video.PublishResponse, error) {
	videoId := idgen.NextId()
	videos := model.Videos{
		BaseModel: model.BaseModel{
			Id: videoId,
		},
		Title:       in.Title,
		VideoUrl:    in.VideoUrl,
		CoverUrl:    in.CoverUrl,
		Description: in.Description,
		AuthorId:    in.AuthorId,
		CommentNum:  0,
		LikeNum:     0,
		TagIds:      nil,
	}
	if err := l.svcCtx.DB.Session(&gorm.Session{}).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Videos{}).Create(&videos).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// 向 feed 服务发送消息
	_, err := l.svcCtx.FeedRPC.PublishContent(l.ctx, &feed.PublishContentRequest{
		UserId:                in.AuthorId,
		VideoId:               videoId,
		VideoCreatorTimestamp: videos.CreatedAt.Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &video.PublishResponse{VideoId: videoId}, nil
}
