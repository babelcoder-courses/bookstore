//go:build mage

package main

import (
	"github.com/babelcoder-courses/bookstore/config"
	"github.com/babelcoder-courses/bookstore/migration"
	"github.com/babelcoder-courses/bookstore/seed"
	"github.com/magefile/mage/mg"
)

func loadConfig() error {
	return config.Load()
}

func initDB() error {
	return config.InitDB()
}

type DB mg.Namespace

func (DB) Migrate() error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.Migrate()
}

func (DB) Rollback() error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.RollbackLast()
}

func (DB) RollbackTo(version string) error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.RollbackTo(version)
}

func (DB) Clean() error {
	mg.SerialDeps(loadConfig, initDB)

	return config.DB().Migrator().DropTable(
		"migrations",
		"users",
		"authors",
		"categories",
		"books",
		"book_users",
		"comments",
	)
}

func (db DB) Seed() error {
	mg.SerialDeps(loadConfig, initDB, db.Clean, db.Migrate)
	seed.Load()

	return nil
}
