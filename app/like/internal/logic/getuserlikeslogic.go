package logic

import (
	"context"
	"errors"
	"strconv"
	"zhihu/app/like/model"

	"github.com/redis/go-redis/v9"

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
	key := model.GetLikeRecordKey("video_like", in.UserId)
	result, _ := l.svcCtx.RDB.ZRangeByScore(l.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Count:  -1,
		Offset: 0,
	}).Result()
	if len(result) > 0 {
		postIds := make([]int64, 0, len(result))
		for _, v := range result {
			if v == "" {
				continue
			}
			if v == "-1" {
				continue
			}
			objId, _ := strconv.ParseInt(v, 10, 64)
			postIds = append(postIds, objId)
		}
		return &like.GetUserLikesResponse{
			PostIds: postIds,
		}, nil
	}
	// 2、查询缓存失败，查询数据库
	posts := make([]*model.LikeRecord, 0)
	if err := l.svcCtx.DB.Model(&model.LikeRecord{}).Find(&posts, &model.LikeRecord{
		BizId:  "video_like",
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
