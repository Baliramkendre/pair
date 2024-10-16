package routes

import (
	handlers "pairs/src/api/handler"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() *Router {
	route := mux.NewRouter()
	r := &Router{
		Router: route,
	}
	mainRoute := route.PathPrefix("/api/").Subrouter()

	mainRoute.HandleFunc("/find-pairs", handlers.FindPairs).Methods("POST")
	return r
}
