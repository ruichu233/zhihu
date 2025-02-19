package main

import (
	"context"
	"encoding/json"
	"time"
	"zhihu/app/user/model"
	"zhihu/pkg/localqueue"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	_rdb *redis.Client
	_db  *gorm.DB
)

var (
	UserQueue = localqueue.NewBatchQueue(10000, 100, 1*time.Second, localqueue.BatchMode)
)

func InitQueue(rdb *redis.Client, db *gorm.DB) {
	_rdb = rdb
	_db = db
	// 初始化全局变量
	go UserQueue.Run(1, UserQueueHandler)
}

func CloseQueue() {
	UserQueue.Stop()
}

func UserQueueHandler(msgs []interface{}) error {
	pipe := _rdb.Pipeline()
	for _, id := range msgs {
		idInt := id.(int64)
		pipe.Get(context.Background(), model.GetUserInfoKey(idInt))
	}
	cmds, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}
	var users []*model.User
	for _, cmd := range cmds {
		res := cmd.(*redis.StringCmd)
		if res.Err() == redis.Nil {
			continue
		}
		if res.Val() == "" {
			continue
		}
		var user model.User
		if err = json.Unmarshal([]byte(res.Val()), &user); err != nil {
			continue
		}
		users = append(users, &user)
	}
	// 更新用户信息
	if len(users) > 0 {
		// 批量更新数据库
		err := _db.Transaction(func(tx *gorm.DB) error {
			for _, user := range users {
				if err := tx.Model(&model.User{}).Where("id = ?", user.Id).
					Updates(map[string]interface{}{
						"follow_count":   user.FollowCount,
						"follower_count": user.FollowerCount,
					}).Error; err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
