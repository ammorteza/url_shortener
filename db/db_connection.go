package db

import "github.com/jinzhu/gorm"

type DbConnection interface {
	Connect() (*gorm.DB, error)
}
