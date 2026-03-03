package controllers

import (
	"encoding/json"
	"gophernet/authentication"
	"gophernet/src/data"
	"gophernet/src/models"
	"gophernet/src/repositories"
	"gophernet/src/responses"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication cria novas publicacoes de usuarios na aplicacao
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publications models.Publications

	if err = json.Unmarshal(bodyRequest, &publications); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	publications.AuthorID = userID

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	publications.ID, err = repo.Create(publications)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publications.ID)
}

func FindPublicationByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewPublicationsRepositories(db)
	publication, err := repo.FindByID(publicationID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

func SearchPublications(w http.ResponseWriter, r *http.Request) {

}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
