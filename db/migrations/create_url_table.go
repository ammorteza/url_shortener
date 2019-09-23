package migrations

import (
	"errors"
	"github.com/ammorteza/url_shortener/db"
	"github.com/ammorteza/url_shortener/models"
)

type UrlTableMigration struct {
	Migration
}

func NewUrlTableMigration(connection db.DbConnection) (MigrationInterface, error){
	urlTableMigration := UrlTableMigration{}
	dbConnection, err := connection.Connect()
	urlTableMigration.db = dbConnection
	var urlTable MigrationInterface = urlTableMigration
	return urlTable, err
}

func (ut UrlTableMigration) Make() error{
	if !ut.db.HasTable(&models.Url{}) {
		return ut.db.CreateTable(&models.Url{}).Error
	}else{
		return errors.New("duplicate error: insert new record in url table!")
	}
}

func (ut UrlTableMigration) Drop() error{
	return ut.db.DropTableIfExists(&models.Url{}).Error
}

func (ut UrlTableMigration) Reset() error{
	return ut.db.DropTableIfExists(&models.Url{}).Error
}
