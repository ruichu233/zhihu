package model

type Notification struct {
	BaseModel
	ToUserId   int64  `gorm:"column:to_user_id;type:bigint(20);not null" json:"to_user_id"`
	FromUserId int64  `gorm:"column:from_user_id;type:bigint(20);not null" json:"from_user_id"`
	Content    string `gorm:"column:content;type:varchar(255);not null" json:"content"`
	Type       int32  `gorm:"column:type;type:tinyint(4);not null" json:"type"` // 1: 关注通知 2: 点赞通知 3: 评论通知
}

func (n *Notification) TableName() string {
	return "notifications"
}