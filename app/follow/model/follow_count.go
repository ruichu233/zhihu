package model

import "zhihu/pkg/model"

type FollowsCount struct {
	model.BaseModel
	UserId       int64 `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	FollowsCount int64 `gorm:"column:follow_count;type:bigint(20);not null" json:"follow_count"`
	FansCount    int64 `gorm:"column:fans_count;type:bigint(20);not null" json:"fans_count"`
}
