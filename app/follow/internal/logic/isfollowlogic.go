package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/model"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowLogic {
	return &IsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowLogic) IsFollow(in *follow.IsFollowRequest) (*follow.IsFollowResponse, error) {
	key := fmt.Sprintf("follow_model:%d", in.UserId)
	mkey := fmt.Sprintf("%d", in.ToUserId)
	res, err := l.svcCtx.RDB.HGet(l.ctx, key, mkey).Result()
	if err != nil {
		return nil, err
	}
	if res != "" {
		var followModel model.Follow
		if err := json.Unmarshal([]byte(res), &followModel); err != nil {
			return nil, err
		}
		return &follow.IsFollowResponse{
			IsFollow: !followModel.DeletedAt.Valid,
		}, nil
	}

	// 从数据库查询
	var followModel model.Follow
	if err := l.svcCtx.DB.Model(&model.Follow{}).Where("follower_id = ? and followee_id = ?", in.UserId, in.ToUserId).First(&followModel).Error; err != nil {
		return nil, err
	}
	if followModel.DeletedAt.Valid {
		return &follow.IsFollowResponse{
			IsFollow: false,
		}, nil
	}

	return &follow.IsFollowResponse{
		IsFollow: true,
	}, nil
}
