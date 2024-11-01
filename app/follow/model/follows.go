package model

type Follow struct {
	BaseModel
	FollowerID int64 `gorm:"column:follower_id;type:bigint(20);not null" json:"follower_id"` // 关注者的ID
	FolloweeID int64 `gorm:"column:followee_id;type:bigint(20);not null" json:"followee_id"` // 被关注者的ID
}

func (Follow) TableName() string {
	return "follows"
}
