package localqueue

import (
	"time"
	"zhihu/app/notification/model"
	"zhihu/pkg/localqueue"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	_rdb        *redis.Client
	_db         *gorm.DB
	NotifyQueue *localqueue.BatchQueue
)

func Init(rdb *redis.Client, db *gorm.DB) {
	_rdb = rdb
	_db = db
	NotifyQueue = localqueue.NewBatchQueue(10000, 100, time.Second*1, localqueue.BatchMode)
	Run()
}

func Run() {
	NotifyQueue.Run(1, func(batch []interface{}) error {
		for _, data := range batch {
			notify := data.(*model.Notification)
			if err := _db.Model(&model.Notification{}).Create(notify).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
