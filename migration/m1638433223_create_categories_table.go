package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638433223CreateCategoriesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638433223",
		Migrate: func(tx *gorm.DB) error {
			type category struct {
				gorm.Model
				Name string `gorm:"unique;not null"`
				Desc string `gorm:"not null"`
			}

			return tx.Migrator().CreateTable(&category{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("category")
		},
	}
}
