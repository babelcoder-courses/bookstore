package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638433781CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638433781",
		Migrate: func(tx *gorm.DB) error {
			type user struct {
				gorm.Model
				Name     string `gorm:"unique;not null"`
				Email    string `gorm:"uniqueIndex; not null"`
				Password string `gorm:"not null"`
				Avatar   string
				Role     uint8 `gorm:"default:3; not null"`
			}

			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("user")
		},
	}
}
