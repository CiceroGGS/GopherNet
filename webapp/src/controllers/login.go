package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/responses"
)

// Login utiliza o email e senha do usuario para autenticar na aplicacao
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIErro{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIErro{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	responses.JSON(w, http.StatusOK, nil)
}
