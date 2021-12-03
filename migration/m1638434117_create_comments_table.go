package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638434117CreateCommentsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638434117",
		Migrate: func(tx *gorm.DB) error {
			type comment struct {
				gorm.Model
				Body   string
				UserID uint
			}

			return tx.Migrator().CreateTable(&comment{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("comment")
		},
	}
}
