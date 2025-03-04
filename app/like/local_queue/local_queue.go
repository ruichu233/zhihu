package local_queue

import (
	"context"
	"fmt"
	"time"
	"zhihu/app/like/model"
	"zhihu/app/like/pb/like"
	"zhihu/app/like/types"
	"zhihu/pkg/localqueue"

	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
)

var (
	_rdb *redis.Client
	_db  *gorm.DB
)

var (
	LikeQueue = localqueue.NewBatchQueue(10000, 100, 1*time.Second, localqueue.BatchMode)
)

func InitQueue(rdb *redis.Client, db *gorm.DB) {
	_rdb = rdb
	_db = db
	// 初始化全局变量
	go LikeQueue.Run(1, LikeQueueHandler)
}

func CloseQueue() {
	LikeQueue.Stop()
}

func LikeQueueHandler(msg []interface{}) error {
	for _, v := range msg {
		value := v.(*types.LikeAction)
		if value.ActionType == like.LikeActionRequest_LIKE {
			addLikeRecord(context.Background(), _rdb, model.GetLikeRecordKey(value.BizId, value.UserId), fmt.Sprintf("%d", value.ObjId), float64(time.Now().Unix()))
			var likeRecord model.LikeRecord
			if err := _db.Model(&model.LikeRecord{}).Unscoped().Where("biz_id = ? and obj_id = ? and user_id = ?", value.BizId, value.ObjId, value.UserId).Limit(1).Find(&likeRecord).Error; err != nil {
				return err
			}
			if likeRecord.Id == 0 {
				_db.Model(&model.LikeRecord{}).Create(&model.LikeRecord{
					BaseModel: model.BaseModel{
						Id: idgen.NextId(),
					},
					BizId:  value.BizId,
					ObjId:  value.ObjId,
					UserId: value.UserId,
				})
			} else {
				_db.Model(&model.LikeRecord{}).Unscoped().Where("biz_id =? and obj_id =? and user_id =?", value.BizId, value.ObjId, value.UserId).Updates(map[string]interface{}{
					"updated_at": time.Now(),
					"deleted_at": nil,
				})
			}
			// 增加点赞数

			_rdb.Incr(context.Background(), model.GetLikeCountKey(value.BizId, value.ObjId))

		} else {
			_db.Model(&model.LikeRecord{}).Where("biz_id =? and obj_id =? and user_id =?", value.BizId, value.ObjId, value.UserId).Delete(&model.LikeRecord{})
			removeLikeRecord(context.Background(), _rdb, model.GetLikeRecordKey(value.BizId, value.UserId), fmt.Sprintf("%d", value.ObjId))
			// 减少点赞数
			_rdb.Decr(context.Background(), model.GetLikeCountKey(value.BizId, value.ObjId))
		}
	}
	return nil
}

// 添加点赞记录并维护前 100 条
func addLikeRecord(ctx context.Context, rdb *redis.Client, key string, member string, score float64) error {
	// 开启 Redis 事务
	pipe := rdb.TxPipeline()
	// 新增点赞记录
	pipe.ZAdd(ctx, "like_records", redis.Z{
		Score:  score,
		Member: member,
	})
	// 获取有序集合的元素数量
	pipe.ZCard(ctx, "like_records")

	// 执行事务
	results, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	// 获取元素数量
	count := results[1].(*redis.IntCmd).Val()
	if count > 100 {
		// 计算需要移除的元素数量
		removeCount := count - 100
		// 移除排名靠后的元素
		_, err := rdb.ZRemRangeByRank(ctx, key, 0, removeCount-1).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func removeLikeRecord(ctx context.Context, rdb *redis.Client, key string, member string) error {
	// 开启 Redis 事务
	pipe := rdb.TxPipeline()
	pipe.ZRem(ctx, key, member)
	// 执行事务
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
