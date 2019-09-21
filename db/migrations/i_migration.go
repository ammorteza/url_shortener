package migrations

import "github.com/ammorteza/url_shortener/db"

type MigrationInterface interface {
	Make()
	Drop()
}

type MigrationCliInterface interface {
	Make(connection db.DbConnection)
	Reset(connection db.DbConnection)
}
