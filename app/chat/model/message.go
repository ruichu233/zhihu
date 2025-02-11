package model

type Message struct {
	BaseModel
	ToUserId   int64  `gorm:"column:to_user_id;type:bigint(20);not null" json:"to_user_id"`
	FromUserId int64  `gorm:"column:from_user_id;type:bigint(20);not null" json:"from_user_id"`
	Content    string `gorm:"column:content;type:varchar(255);not null" json:"content"`
}

func (m *Message) TableName() string {
	return "message"
}
