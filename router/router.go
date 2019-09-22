package router

import (
	"fmt"
	"github.com/ammorteza/url_shortener/controllers"
	"github.com/ammorteza/url_shortener/db"
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
	dbConnection db.DbConnection
}

func New(connection db.DbConnection) *Router{
	router := &Router{}
	router.dbConnection = connection
	urlController := controllers.NewUrlController(connection)
	mainController := controllers.NewMainController(connection)

	router.routes = []Route{
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
	return router
}

func (r *Router)buildRoute()  {
	for _, route := range r.routes {
		r.router.HandleFunc(route.path, route.handler).Methods(route.method)
		fmt.Println(route)
	}
}

func (r *Router) Start() {
	r.router = mux.NewRouter().StrictSlash(true)
	r.buildRoute()
	fmt.Println("start http connection at 8080 tcp port")
	log.Fatal(http.ListenAndServe(":8080", r.router))
}
