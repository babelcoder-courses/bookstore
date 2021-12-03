package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name   string `gorm:"unique;not null"`
	Bio    string
	Avatar string
	Books  []Book
}
