package migration

import (
	"github.com/babelcoder-courses/bookstore/config"
	"github.com/go-gormigrate/gormigrate/v2"
)

func newMigrate() *gormigrate.Gormigrate {
	db := config.DB()

	return gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1638433223CreateCategoriesTable(),
			m1638433494CreateAuthorsTable(),
			m1638433601CreateBooksTable(),
			m1638433781CreateUsersTable(),
			m1638434031CreateBookUsersTable(),
			m1638434117CreateCommentsTable(),
		},
	)
}

func Migrate() error {
	return newMigrate().Migrate()
}

func RollbackLast() error {
	return newMigrate().RollbackLast()
}

func RollbackTo(version string) error {
	return newMigrate().RollbackTo(version)
}
