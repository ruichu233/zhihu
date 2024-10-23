package model

import "zhihu/pkg/model"

type Follows struct {
	model.BaseModel
	UserId         int64 `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	FollowedUserId int64 `gorm:"column:followed_user_id;type:bigint(20);not null" json:"followed_user_id"`
	FollowStatus   int   `gorm:"column:follow_status;type:tinyint(1);not null" json:"follow_status"`
}

func (Follows) TableName() string {
	return "follows"
}
