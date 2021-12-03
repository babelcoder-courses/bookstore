package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638434031CreateBookUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638434031",
		Migrate: func(tx *gorm.DB) error {
			type bookUser struct {
				gorm.Model
				BookID uint
				UserID uint
			}

			return tx.Migrator().CreateTable(&bookUser{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("bookUser")
		},
	}
}
