package main

import (
	"url-shortener/src/db"
	"url-shortener/src/router"
)

func main()  {
	_router := router.Router{}
	_router.Start(db.MysqlConnection{})
}
