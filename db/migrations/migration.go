package migrations

import (
	"github.com/jinzhu/gorm"
)

type Migration struct {
	db *gorm.DB
}
