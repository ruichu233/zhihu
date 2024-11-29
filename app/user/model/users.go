package model

type User struct {
	BaseModel
	Username string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Avatar   string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
}

func (User) TableName() string {
	return "users"
}
