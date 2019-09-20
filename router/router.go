package router

import (
	"fmt"
	"github.com/ammorteza/urlShortener/controllers"
	"github.com/ammorteza/urlShortener/interfaces"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Route struct {
	path    string
	method  string
	handler func(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	router       *mux.Router
	routes       []Route
	dbConnection interfaces.DbConnection
}

func (r *Router) makeRoutes() {
	var urlController controllers.UrlController
	urlController.Init(r.dbConnection)

	var mainController controllers.MainController
	mainController.Init(r.dbConnection)

	r.routes = []Route{
		{
			path:    "/{unique_url_id}",
			method:  "GET",
			handler: urlController.RedirectShortenUrl,
		},
		{
			path:    "/url/shorten/make",
			method:  "POST",
			handler: urlController.MakeUrl,
		},
		{
			path:    "/page/hamedan/trip/hotel/get_all",
			method:  "GET",
			handler: mainController.MainUrl,
		},
	}

	for _, route := range r.routes {
		r.router.HandleFunc(route.path, route.handler).Methods(route.method)
		fmt.Println(route)
	}
}

func (r *Router) setDbConnection(dbConnection interfaces.DbConnection) {
	r.dbConnection = dbConnection
}

func (r *Router) Start(dbConnection interfaces.DbConnection) {
	r.setDbConnection(dbConnection)
	r.router = mux.NewRouter().StrictSlash(true)
	r.makeRoutes()
	fmt.Println("start http connection at 8080 tcp port")
	log.Fatal(http.ListenAndServe(":8080", r.router))
}
