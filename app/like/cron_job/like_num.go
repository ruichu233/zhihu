package cronjob

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"zhihu/app/like/model"

	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var tiker *time.Ticker
var _rdb *redis.Client
var _db *gorm.DB

var script = `
local keys = redis.call('KEYS', ARGV[1])
local result = {}
for i, key in ipairs(keys) do
    local value = redis.call('GET', key)
    result[i] = {key, value}
end
return cjson.encode(result)
`

func Init(rdb *redis.Client, db *gorm.DB) {
	tiker = time.NewTicker(1 * time.Minute)
	_rdb = rdb
	_db = db
	go Run()
}
func Run() {
	for {
		select {
		case <-tiker.C:
			// 执行定时任务
			// 获取点赞数
			result, err := _rdb.Eval(context.Background(), script, []string{}, "count:patten:*").Result()
			if err != nil {
				logx.Errorf("redis.Eval error: %v", err)
				continue
			}
			var resultSlice [][]string
			if err := json.Unmarshal([]byte(result.(string)), &resultSlice); err != nil {
				logx.Errorf("json.Unmarshal error: %v", err)
				continue
			}
			for _, item := range resultSlice {
				key := item[0]
				value := item[1]
				var bizId string
				var objId int64
				var likeNum int64
				if value == "-1" {
					continue
				} else {
					keySlice := strings.Split(key, ":")
					bizId = keySlice[2]
					objId, err = strconv.ParseInt(keySlice[3], 10, 64)
					if err != nil {
						logx.Errorf("strconv.ParseInt error: %v", err)
						continue
					}
					likeNum, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						logx.Errorf("strconv.ParseInt error: %v", err)
						continue
					}
				}
				var likeCount model.LikeCount
				if err := _db.Model(&model.LikeCount{}).Where("biz_id =? and obj_id =?", bizId, objId).Limit(1).Find(&likeCount).Error; err != nil {
					logx.Errorf("db.Model error: %v", err)
					continue
				}
				if likeCount.Id == 0 {
					_db.Create(&model.LikeCount{
						BaseModel: model.BaseModel{
							Id: idgen.NextId(),
						},
						BizId:   bizId,
						ObjId:   objId,
						LikeNum: likeNum,
					})
				} else {
					_db.Model(&model.LikeCount{}).Where("biz_id = ? and obj_id = ?", bizId, objId).Update("like_num", likeNum)
				}
				tiker.Reset(1 * time.Minute)
			}
		}
	}
}
