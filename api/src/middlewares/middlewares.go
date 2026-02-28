package middlewares

import (
	"gophernet/authentication"
	"gophernet/src/responses"
	"log"
	"net/http"
)

// Logger escreve informacoes da requisicao no log
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n###################################################\n - [ %s ] - [ %s ] - [ %s ] - \n###################################################", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate verifica se o usuario fazendo a requisicao esta autenticado
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.TokenValidation(r); err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
