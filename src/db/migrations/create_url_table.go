package migrations

import (
	"url-shortener/src/models"
)

type UrlTableMigration struct {
	Migration
}

func (ut UrlTableMigration)Make()  {
	if !ut.db.HasTable(&models.Url{}){
		ut.db.CreateTable(&models.Url{})
	}
}

func (ut UrlTableMigration)Drop()  {
	ut.db.DropTableIfExists(&models.Url{})
}
