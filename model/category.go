package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Desc  string `gorm:"not null"`
	Books []Book
}
