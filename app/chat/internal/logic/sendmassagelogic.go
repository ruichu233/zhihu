package logic

import (
	"context"

	"zhihu/app/chat/internal/svc"
	"zhihu/app/chat/model"
	"zhihu/app/chat/pb/chat"

	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendMassageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMassageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMassageLogic {
	return &SendMassageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMassageLogic) SendMassage(in *chat.SendMassageRequest) (*chat.SendMassageResponse, error) {
	id := idgen.NextId()
	// 存入数据库
	if err := l.svcCtx.DB.Model(&model.Message{}).Create(&model.Message{
		BaseModel: model.BaseModel{
			Id: id,
		},
		FromUserId: in.SenderId,
		ToUserId:   in.ReceiverId,
		Content:    in.Content,
	}).Error; err != nil {
		return nil, err
	}

	return &chat.SendMassageResponse{
		Id: id,
	}, nil
}
