package migrations

import (
	"fmt"
	"github.com/ammorteza/urlShortener/interfaces"
)

type MigrationCli struct {
}

func (m MigrationCli) Make(connection interfaces.DbConnection) {
	fmt.Println("Migrate all database tables start")
	////////////////// make url table migration ///////////////////
	fmt.Println("Migrate urlTableMigration file ...")
	UrlTableMigration := UrlTableMigration{}
	UrlTableMigration.Init(connection)
	var utmInterface interfaces.MigrationInterface = UrlTableMigration
	utmInterface.Make()
	///////////////////////////////////////////////////////////////
}

func (m MigrationCli) Reset(connection interfaces.DbConnection) {
	fmt.Println("Resetting all database tables start")
	////////////////// make url table migration ///////////////////
	fmt.Println("Reset urlTable ...")
	UrlTableMigration := UrlTableMigration{}
	UrlTableMigration.Init(connection)
	var utmInterface interfaces.MigrationInterface = UrlTableMigration
	utmInterface.Drop()
	///////////////////////////////////////////////////////////////
}
