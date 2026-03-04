package controllers

import (
	"encoding/json"
	"errors"
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

// SearchPublications retorna todas as publicacoes de todos usuarios que o usuario da aplicao segue
func SearchPublications(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewPublicationsRepositories(db)
	publications, err := repo.Search(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

// UpdatePublication permite que um usuário autenticado atualize uma publicação de sua própria autoria.
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	publicationInDB, err := repo.FindByID(publicationID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicationInDB.AuthorID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Nao e possivel atualizar a publicacao de outro usuario"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var publication models.Publications

	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// if err = publication.Prepare(); err != nil {
	// 	responses.Erro(w, http.StatusBadRequest, err)
	// 	return
	// }

	if err = repo.Update(publicationID, publication); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

// DeletePublication permite que um usuário autenticado exclua uma publicação de sua própria autoria.
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	publicationInDB, err := repo.FindByID(publicationID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if publicationInDB.AuthorID != userID {
		responses.Erro(w, http.StatusBadRequest, errors.New("Nao e possivel deletar a publicacoa de outro usuario"))
		return
	}

	if err = repo.Delete(publicationID); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

// SearchPublicationsByUser retorna todas publicacoes de um usuario especifico
func SearchPublicationsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	publicactions, err := repo.SearchByUser(userID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicactions)
}

// LikePublication adiciona uma curtina na publicacao
func LikePublication(w http.ResponseWriter, r *http.Request) {
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
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	if err = repo.LikePublication(publicationID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

// DislikePublication remove uma curtina na publicacao
func DislikePublication(w http.ResponseWriter, r *http.Request) {
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
	defer db.Close()

	repo := repositories.NewPublicationsRepositories(db)
	if err = repo.DislkePublication(publicationID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
