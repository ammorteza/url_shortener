package main

import (
	"github.com/ammorteza/urlShortener/src/db"
	"github.com/ammorteza/urlShortener/src/router"
)

func main()  {
	_router := router.Router{}
	_router.Start(db.MysqlConnection{})
}
