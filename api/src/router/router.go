package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar cria a instância principal do roteador.
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
