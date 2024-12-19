package logic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLikersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLikersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLikersLogic {
	return &GetPostLikersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询点赞人列表
func (l *GetPostLikersLogic) GetPostLikers(in *like.GetPostLikersRequest) (*like.GetPostLikersResponse, error) {
	key := fmt.Sprintf("%s:likes:%d", in.BizId, in.ObjId)
	// 1、查询缓存
	result, err := l.svcCtx.RDB.ZRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Count:  10,
		Offset: 0,
	}).Result()
	if err == nil {
		userIds := make([]int64, 0, len(result))
		for _, v := range result {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			userIds = append(userIds, id)
		}
		return &like.GetPostLikersResponse{
			UserIds: userIds,
		}, nil
	}
	// 2、查询数据库
	return nil, err
}
