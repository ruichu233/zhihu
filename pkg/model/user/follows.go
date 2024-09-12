package user

import "zhihu/pkg/model"

type Follows struct {
	model.BaseModel
	FollowerId int64 `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`     // 关注者id
	FollowedId int64 `gorm:"column:follow_id;type:bigint(20);not null" json:"follow_id"` // 被关注者id
}

func (Follows) TableName() string {
	return "follows"
}
