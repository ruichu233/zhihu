package model

type Comments struct {
	BaseModel
	BizId         string `gorm:"column:biz_id;type:varchar(50);not null" json:"biz_id"`
	ObjId         int64  `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	ReplyUserId   int64  `gorm:"column:reply_user_id;type:bigint(20);not null" json:"reply_user_id"`
	BeReplyUserId int64  `gorm:"column:be_reply_user_id;type:bigint(20);not null" json:"be_reply_user_id"`
	ParentId      int64  `gorm:"column:parent_id;type:bigint(20);not null" json:"parent_id"`
	Content       string `gorm:"column:content;type:varchar(255);not null" json:"content"`
	Status        uint8  `gorm:"column:status;type:tinyint(1);not null" json:"status"`
	LikeNum       int64  `gorm:"column:like_num;type:bigint(20);not null" json:"like_num"`
}

func (c *Comments) TableName() string {
	return "comments"
}

type CommentCount struct {
	BaseModel
	BizId          string `gorm:"column:biz_id;type:varchar(50);not null" json:"biz_id"`
	ObjId          int64  `gorm:"column:obj_id;type:bigint(20);not null" json:"obj_id"`
	CommentNum     int64  `gorm:"column:comment_num;type:bigint(20);not null" json:"comment_num"`
	CommentRootNum int64  `gorm:"column:comment_root_num;type:bigint(20);not null" json:"comment_root_num"`
}

func (c *CommentCount) TableName() string {
	return "comment_count"
}
