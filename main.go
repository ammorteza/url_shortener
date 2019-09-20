package main

import (
	"github.com/ammorteza/urlShortener/db"
	"github.com/ammorteza/urlShortener/router"
)

func main() {
	_router := router.Router{}
	_router.Start(db.MysqlConnection{})
}
