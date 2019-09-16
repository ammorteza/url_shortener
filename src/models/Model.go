package models

import (
	"github.com/jinzhu/gorm"
	"url-shortener/src/interfaces"
)

type Model struct {
	db *gorm.DB
}

func (m *Model)Init(dbConnection interfaces.DbConnection) {
	db := dbConnection.Connect()
	m.db = db
}