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

func (mt *MigrateTables)Make() error{
	fmt.Println("start tables migration")
	//////////// url model /////////////////
	urlTable, err := NewUrlTableMigration(mt.dbConnection)
	if err != nil{
		return err
	}
	if err := urlTable.Make(); err != nil{
		return err
	}
	////////////////////////////////////////
	fmt.Println("finished tables migration")
	return nil
}

func (mt *MigrateTables)Reset() error{
	fmt.Println("start tables reset")
	//////////// url model /////////////////
	urlTable,err := NewUrlTableMigration(mt.dbConnection)
	if err != nil{
		return err
	}
	if err := urlTable.Reset(); err != nil{
		return err
	}
	////////////////////////////////////////
	fmt.Println("finished tables reset")
	return nil
}