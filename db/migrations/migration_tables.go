package migrations

import (
	"fmt"
	"github.com/ammorteza/url_shortener/db"
)

type MigrateTables struct {
	dbConnection db.DbConnection
}

func NewMigrate(connection db.DbConnection) *MigrateTables{
	migrateTable := MigrateTables{}
	migrateTable.dbConnection = connection
	return &migrateTable
}

func (mt *MigrateTables)Make() {
	fmt.Println("start tables migration")
	//////////// url model /////////////////
	urlTable := NewUrlTableMigration(mt.dbConnection)
	urlTable.Make()
	////////////////////////////////////////
	fmt.Println("finished tables migration")
}

func (mt *MigrateTables)Reset() {
	fmt.Println("start tables reset")
	//////////// url model /////////////////
	urlTable := NewUrlTableMigration(mt.dbConnection)
	urlTable.Reset()
	////////////////////////////////////////
	fmt.Println("finished tables reset")
}