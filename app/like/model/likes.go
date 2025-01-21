package model

import (
	"fmt"
)

type LikeRecord struct {
	BaseModel
	BizId  string `gorm:"column:biz_id;type:varchar(50);not null" json:"biz_id"`
	ObjId  int64  `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	UserId int64  `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
}

func (l *LikeRecord) TableName() string {
	return "like_records"
}

type LikeCount struct {
	BaseModel
	BizId   string `gorm:"column:biz_id;type:varchar(50);not null" json:"biz_id"`
	ObjId   int64  `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	LikeNum int64  `gorm:"column:like_num;type:bigint(20);not null" json:"like_num"`
}

func (l *LikeCount) TableName() string {
	return "like_counts"
}

func GetLikeRecordKey(bizId string, userId int64) string {
	return fmt.Sprintf("user:likes:patten:%d:%s", userId, bizId)
}

func GetLikeCountKey(bizId string, objId int64) string {
	return fmt.Sprintf("count:patten:%d:%s", objId, bizId)
}
