package models

import (
	"github.com/ammorteza/urlShortener/interfaces"
	"github.com/jinzhu/gorm"
)

type Model struct {
	db *gorm.DB
}

func (m *Model) Init(dbConnection interfaces.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}
