package user

import "zhihu/pkg/model"

type Users struct {
	model.BaseModel
	UserName string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
}

func (Users) TableName() string {
	return "users"
}
