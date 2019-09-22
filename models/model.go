package models

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	db *gorm.DB
}
