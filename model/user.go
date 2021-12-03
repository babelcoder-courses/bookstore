package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Email    string `gorm:"uniqueIndex; not null"`
	Password string `gorm:"not null"`
	Avatar   string
	Role     Role    `gorm:"default:3; not null"`
	Books    []*Book `gorm:"many2many:book_users"`
}
