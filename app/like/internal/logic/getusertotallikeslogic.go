package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"zhihu/app/like/model"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTotalLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTotalLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTotalLikesLogic {
	return &GetUserTotalLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户收到的总点赞数
func (l *GetUserTotalLikesLogic) GetUserTotalLikes(in *like.GetUserTotalLikesRequest) (*like.GetUserTotalLikesResponse, error) {
	// 1、 在缓存中查询点赞总数
	count, err := l.svcCtx.RDB.Get(l.ctx, GetUserTotalLikesKey(in.UserId)).Int64()
	if err == nil {
		return &like.GetUserTotalLikesResponse{
			TotalLikes: count,
		}, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}
	// 2、缓存中没有查询数据，查询数据库
	if err := l.svcCtx.DB.Model(&model.LikeRecord{}).Where("user_id = ?", in.UserId).Count(&count).Error; err != nil {
		return nil, errors.New("查询点赞总数失败")
	}
	// 3、更新缓存
	l.svcCtx.RDB.Set(l.ctx, GetUserTotalLikesKey(in.UserId), count, 0)

	return &like.GetUserTotalLikesResponse{
		TotalLikes: count,
	}, nil
}

func GetUserTotalLikesKey(userId int64) string {
	return fmt.Sprintf("USER_TOTAL_LIKES_%d", userId)
}
