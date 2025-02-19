package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"zhihu/app/notification/internal/svc"
	"zhihu/app/notification/model"
	"zhihu/app/notification/pb/notification"
)

type GetNotificationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotificationsLogic {
	return &GetNotificationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNotificationsLogic) GetNotifications(in *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	var notifications []model.Notification
	offset := (in.Page - 1) * in.PageSize

	err := l.svcCtx.DB.Where("to_user_id = ?", in.UserId).Order("created_at desc").Offset(int(offset)).Limit(int(in.PageSize)).Find(&notifications).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var resp notification.NotificationResponse
	resp.StatusCode = 0
	resp.StatusMsg = "success"

	for _, n := range notifications {
		resp.Notifications = append(resp.Notifications, &notification.Notification{
			Id:         n.ID,
			ToUserId:   n.ToUserId,
			FromUserId: n.FromUserId,
			Content:    n.Content,
			Type:       n.Type,
			CreatedAt:  n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &resp, nil
}