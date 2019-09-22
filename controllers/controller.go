package controllers

import (
	"github.com/ammorteza/url_shortener/db"
)

type Controller struct {
	dbConnection db.DbConnection
}
