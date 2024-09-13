package comment

import "zhihu/pkg/model"

type Comment struct {
	model.BaseModel
	Pid     int64  `gorm:"column:pid;type:bigint(20);not null" json:"pid"`
	Content string `gorm:"column:content;type:varchar(255);not null" json:"content"`
}
