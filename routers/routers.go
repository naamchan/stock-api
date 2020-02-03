package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naamchan/routers/routes"
)

// RequestHandler is an interface for every request handler
type RequestHandler interface {
	GetRoute() string
	HandleRequest(http.ResponseWriter, *http.Request)
}

// Init router
func Init() {
	router := mux.NewRouter().StrictSlash(true)

	addRoute(router, routes.NewIndexHandler())
}

func addRoute(router *mux.Router, requestHandler RequestHandler) {
	router.HandleFunc(requestHandler.GetRoute(), func(w http.ResponseWriter, r *http.Request) {
		requestHandler.HandleRequest(w, r)
	})
}
