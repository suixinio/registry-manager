package model

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Name     string `gorm:"size:50;uniqueIndex:idx_name"`
	Password string `json:"-"`
	Email    string `gorm:"type:varchar(100)"`
	Status   int
	Authn    string `gorm:"type:text"`
}
