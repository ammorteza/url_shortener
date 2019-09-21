package controllers

import (
	"github.com/ammorteza/url_shortener/db"
)

type Controller struct {
	dbConnection db.DbConnection
}

func (c *Controller) Init(dbConnect db.DbConnection) {
	c.dbConnection = dbConnect
}
