package logic

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"zhihu/app/like/model"

	"zhihu/app/like/internal/svc"
	"zhihu/app/like/pb/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikesLogic {
	return &GetUserLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询某个用户的点赞列表
func (l *GetUserLikesLogic) GetUserLikes(in *like.GetUserLikesRequest) (*like.GetUserLikesResponse, error) {
	// 1、查询缓存
	key := model.GetLikeRecordKey("video", in.UserId)
	result, err := l.svcCtx.RDB.ZRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Count:  -1,
		Offset: 0,
	}).Result()
	if err == nil {
		postIds := make([]int64, 0, len(result))
		for _, v := range result {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			postIds = append(postIds, id)
		}
		return &like.GetUserLikesResponse{
			PostIds: postIds,
		}, nil
	}
	if !errors.Is(err, redis.Nil) {
		return nil, err
	}
	// 2、查询缓存失败，查询数据库
	posts := make([]*model.LikeRecord, 0)
	if err := l.svcCtx.DB.Model(&model.LikeRecord{}).Find(&posts, &model.LikeRecord{
		BizId:  "video",
		UserId: in.UserId,
	}); err != nil {
		return nil, errors.New("查询点赞列表失败")
	}
	postIds := make([]int64, 0, len(posts))
	for _, v := range posts {
		postIds = append(postIds, v.ObjId)
	}
	// 3、更新缓存
	go func() {
		for _, v := range posts {
			l.svcCtx.RDB.ZAdd(l.ctx, key, redis.Z{
				Score:  float64(v.UpdatedAt.Unix()),
				Member: v.ObjId,
			})
		}
	}()

	return &like.GetUserLikesResponse{
		PostIds: postIds,
	}, nil
}
