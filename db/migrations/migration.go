package migrations

import (
	"github.com/ammorteza/url_shortener/interfaces"
	"github.com/jinzhu/gorm"
)

type Migration struct {
	db *gorm.DB
}

func (m *Migration) Init(dbConnection interfaces.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}
