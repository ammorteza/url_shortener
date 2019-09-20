package controllers

import "github.com/ammorteza/urlShortener/interfaces"

type Controller struct {
	dbConnection interfaces.DbConnection
}

func (c *Controller)Init(dbConnect interfaces.DbConnection){
	c.dbConnection = dbConnect
}
