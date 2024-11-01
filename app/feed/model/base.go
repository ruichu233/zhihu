package model

import "gorm.io/gorm"

type BaseModel struct {
	Id        int64          `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
