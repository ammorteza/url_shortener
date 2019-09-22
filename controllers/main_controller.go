package controllers

import (
	"github.com/ammorteza/url_shortener/db"
	"net/http"
)

type MainController struct {
	Controller
}

func NewMainController(connection db.DbConnection)  *MainController{
	mainController := &MainController{}
	mainController.dbConnection = connection
	return mainController
}

func (c *MainController) MainUrl(w http.ResponseWriter, r *http.Request) {

}
