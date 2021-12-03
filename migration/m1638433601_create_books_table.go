package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638433601CreateBooksTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638433601",
		Migrate: func(tx *gorm.DB) error {
			type book struct {
				gorm.Model
				Title      string `gorm:"unique;not null"`
				Desc       string `gorm:"not null"`
				Image      string `gorm:"not null"`
				CategoryID uint
				AuthorID   uint
			}

			return tx.Migrator().CreateTable(&book{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("book")
		},
	}
}
