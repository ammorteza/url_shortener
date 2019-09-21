package models

import (
	"github.com/ammorteza/url_shortener/db"
	"github.com/jinzhu/gorm"
)

type Model struct {
	db *gorm.DB
}

func (m *Model) Init(dbConnection db.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}
