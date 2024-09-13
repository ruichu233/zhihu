package video

import "zhihu/pkg/model"

type Video struct {
	model.BaseModel
	Title    string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	VideoUrl string `gorm:"column:url;type:varchar(255);not null" json:"url"`
	CoverUrl string `gorm:"column:cover_url;type:varchar(255);not null" json:"cover_url"`
}

func (Video) TableName() string {
	return "video"
}
