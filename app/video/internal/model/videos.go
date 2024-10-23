package video

import "zhihu/pkg/model"

type Videos struct {
	model.BaseModel
	Title       string  `gorm:"column:title;type:varchar(255);not null" json:"title"`
	VideoUrl    string  `gorm:"column:url;type:varchar(255);not null" json:"url"`
	CoverUrl    string  `gorm:"column:cover_url;type:varchar(255);not null" json:"cover_url"`
	Description string  `gorm:"column:description;type:varchar(255);not null" json:"description"`
	AuthorId    int64   `gorm:"column:author_id;type:bigint(20);not null" json:"author_id"`
	CommentNum  int64   `gorm:"column:comment_num;type:bigint(20);not null" json:"comment_num"`
	LikeNum     int64   `gorm:"column:like_num;type:bigint(20);not null" json:"like_num"`
	TagIds      []int64 `gorm:"-" json:"tag_ids"`
}

func (Videos) TableName() string {
	return "videos"
}
