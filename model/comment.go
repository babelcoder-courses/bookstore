package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string
	UserID uint
	User   User
}
