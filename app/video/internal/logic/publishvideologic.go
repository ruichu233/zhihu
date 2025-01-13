package logic

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"zhihu/app/video/internal/model"

	"zhihu/app/video/internal/svc"
	"zhihu/app/video/pb/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布视频
func (l *PublishVideoLogic) PublishVideo(in *video.PublishRequest) (*video.PublishResponse, error) {
	resp := &video.PublishResponse{}
	videoId := idgen.NextId()
	videos := model.Video{
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
	}
	if err := l.svcCtx.DB.Session(&gorm.Session{}).Transaction(func(tx *gorm.DB) error {
		// 保存视频信息
		if err := tx.Model(&model.Video{}).Create(&videos).Error; err != nil {
			return err
		}
		//// 向 feed 服务发送消息
		//_, err := l.svcCtx.FeedRPC.PublishContent(l.ctx, &feed.PublishContentRequest{
		//	UserId:                in.AuthorId,
		//	VideoId:               videoId,
		//	VideoCreatorTimestamp: videos.CreatedAt.Unix(),
		//})
		//if err != nil {
		//	return err
		//}
		return nil
	}); err != nil {
		return nil, err
	}
	resp.VideoId = videoId
	return resp, nil
}
