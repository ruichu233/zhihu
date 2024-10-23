package model

type LikeRecords struct {
	BaseModel
	BizId  int64 `gorm:"column:biz_id;type:bigint(20);not null" json:"biz_id"`
	ObjId  int64 `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	UserId int64 `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	Status uint8 `gorm:"column:status;type:tinyint(1);not null" json:"status"`
}

func (LikeRecords) TableName() string {
	return "like_records"
}

type LikeCount struct {
	BaseModel
	BizId   int64 `gorm:"column:biz_id;type:bigint(20);not null" json:"biz_id"`
	ObjId   int64 `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	LikeNum int64 `gorm:"column:like_num;type:bigint(20);not null" json:"like_num"`
}
