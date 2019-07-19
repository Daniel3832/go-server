package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	fmt.Println(routes)
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method). // GET
			Path(route.Path).      // '/'
			Name(route.Name).      // 'Index'
			Handler(handler)       // Index()
	}
	return router
}
