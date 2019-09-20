package interfaces

import "github.com/jinzhu/gorm"

type DbConnection interface {
	Connect() *gorm.DB
}
