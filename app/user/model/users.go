package model

import "strconv"

type User struct {
	BaseModel
	Username      string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Avatar        string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Password      string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Email         string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	Signature     string `gorm:"column:signature;type:varchar(255)" json:"signature"`
	FollowCount   int64  `gorm:"column:follow_count;type:bigint;not null;default:0" json:"follow_count"`     // 关注数
	FollowerCount int64  `gorm:"column:follower_count;type:bigint;not null;default:0" json:"follower_count"` // 粉丝数
}

func (User) TableName() string {
	return "users"
}

func GetUserInfoKey(userId int64) string {
	return "user:info:" + strconv.FormatInt(userId, 10)
}
