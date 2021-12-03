package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1638433494CreateAuthorsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1638433494",
		Migrate: func(tx *gorm.DB) error {
			type author struct {
				gorm.Model
				Name   string `gorm:"unique;not null"`
				Bio    string
				Avatar string
			}

			return tx.Migrator().CreateTable(&author{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("author")
		},
	}
}
