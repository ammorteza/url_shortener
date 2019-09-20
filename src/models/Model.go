package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ammorteza/urlShortener/src/interfaces"
)

type Model struct {
	db *gorm.DB
}

func (m *Model)Init(dbConnection interfaces.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}