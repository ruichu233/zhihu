package logic

import (
	"context"
	"fmt"
	client "github.com/gorse-io/gorse-go"
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
		//	VideoDescription:      videos.Description,
		//})
		//if err != nil {
		//	return err
		//}
		// 向gorse服务发送消息，新建item
		rowAffected, err := l.svcCtx.Gorse.InsertItem(l.ctx, client.Item{
			ItemId:     fmt.Sprintf("%d", videoId),
			Comment:    videos.Description,
			IsHidden:   false,
			Categories: []string{},
			Timestamp:  videos.CreatedAt.Format("2006-01-02 15:04:05"),
			Labels:     []string{},
		})
		if err != nil {
			return err
		}
		if rowAffected.RowAffected <= 0 {
			logx.Errorf("用户 %d 推送新内容到gorse失败", in.AuthorId)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	resp.VideoId = videoId
	return resp, nil
}
