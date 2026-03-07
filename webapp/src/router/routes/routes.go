package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da aplicacao web
type Route struct {
	URI               string
	Method            string
	Function          func(w http.ResponseWriter, r *http.Request)
	reqAuthentication bool
}

// RouterConfig coloca todas as rodas dentro do router
func RouterConfig(r *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
