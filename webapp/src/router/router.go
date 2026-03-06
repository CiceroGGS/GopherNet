package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorna um router com todas as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.RouterConfig(r)
}
