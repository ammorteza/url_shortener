package migrations

import (
	"github.com/ammorteza/url_shortener/db"
	"github.com/jinzhu/gorm"
)

type Migration struct {
	db *gorm.DB
}

func (m *Migration) Init(dbConnection db.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}
