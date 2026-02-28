package routes

import (
	"gophernet/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes representa todas as rotas da API
type Routes struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// RouterConfig itera sobre as rotas e atribui as configuracoes das rotas dentro do router
func RouterConfig(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		}
		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return r
}
