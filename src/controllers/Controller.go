package controllers

import "url-shortener/src/interfaces"

type Controller struct {
	dbConnection interfaces.DbConnection
}

func (c *Controller)Init(dbConnect interfaces.DbConnection){
	c.dbConnection = dbConnect
}
