package db

import _ "github.com/jinzhu/gorm/dialects/mysql"
import "github.com/jinzhu/gorm"

type MysqlConnection struct {
}

func (mc MysqlConnection) Connect() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:13656195@tcp(urlshortener_db)/url_shortener_db?charset=utf8&parseTime=True&loc=Local")
}
