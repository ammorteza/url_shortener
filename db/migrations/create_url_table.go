package migrations

import (
	"github.com/ammorteza/url_shortener/db"
	"github.com/ammorteza/url_shortener/models"
)

type UrlTableMigration struct {
	Migration
}

func NewUrlTableMigration(connection db.DbConnection) MigrationInterface{
	urlTableMigration := UrlTableMigration{}
	urlTableMigration.db = connection.Connect()
	var urlTable MigrationInterface = urlTableMigration
	return urlTable
}

func (ut UrlTableMigration) Make() {
	if !ut.db.HasTable(&models.Url{}) {
		ut.db.CreateTable(&models.Url{})
	}
}

func (ut UrlTableMigration) Drop() {
	ut.db.DropTableIfExists(&models.Url{})
}

func (ut UrlTableMigration) Reset() {
	ut.db.DropTableIfExists(&models.Url{})
}
