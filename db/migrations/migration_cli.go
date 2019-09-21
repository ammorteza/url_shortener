package migrations

import (
	"fmt"
	"github.com/ammorteza/url_shortener/db"
)

type MigrationCli struct {
}

func (m MigrationCli) Make(connection db.DbConnection) {
	fmt.Println("Migrate all database tables start")
	////////////////// make url table migration ///////////////////
	fmt.Println("Migrate urlTableMigration file ...")
	UrlTableMigration := UrlTableMigration{}
	UrlTableMigration.Init(connection)
	var utmInterface MigrationInterface = UrlTableMigration
	utmInterface.Make()
	///////////////////////////////////////////////////////////////
}

func (m MigrationCli) Reset(connection db.DbConnection) {
	fmt.Println("Resetting all database tables start")
	////////////////// make url table migration ///////////////////
	fmt.Println("Reset urlTable ...")
	UrlTableMigration := UrlTableMigration{}
	UrlTableMigration.Init(connection)
	var utmInterface MigrationInterface = UrlTableMigration
	utmInterface.Drop()
	///////////////////////////////////////////////////////////////
}
