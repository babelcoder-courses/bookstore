package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title      string `gorm:"unique;not null"`
	Desc       string `gorm:"not null"`
	Image      string `gorm:"not null"`
	CategoryID uint
	Category   Category
	AuthorID   uint
	Author     Author
	User       []*User `gorm:"many2many:book_users"`
}
