package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/ammorteza/urlShortener/interfaces"
)

type Migration struct {
	db	*gorm.DB
}

func (m *Migration)Init(dbConnection interfaces.DbConnection){
	db := dbConnection.Connect()
	m.db = db
}
