package controllers

import (
	"encoding/json"
	"fmt"
	"gophernet/authentication"
	"gophernet/src/data"
	"gophernet/src/models"
	"gophernet/src/repositories"
	"gophernet/src/responses"
	"gophernet/src/security"
	"io"
	"net/http"
)

// Login e responsavel pelo login e autenticacao do usuario na API
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Connect()
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	userInDB, err := repo.SearchByEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = security.PasswordVerify(string(userInDB.Password), user.Password); err != nil {
		responses.JSON(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userInDB.ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)
	w.Write([]byte("Logado com sucesso"))
}
