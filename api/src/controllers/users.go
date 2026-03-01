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
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser lê o corpo da requisição, valida e prepara os dados do novo usuário e o persiste no banco de dados.
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare("register"); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers busca todos os usuários no banco de dados filtrando por nome ou nick informado como parâmetro de query na URL (?name=).
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("name"))

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	users, err := repo.Search(nameOrNick)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUFindUserByIDserByID busca um único usuário no banco de dados pelo ID informado na URL.
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repo := repositories.NewUsersRepositories(db)
	user, err := repo.FindByID(ID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser atualiza os dados de um usuário no banco de dados com base no ID informado na URL e nos dados enviados no corpo da requisição.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != ID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é permitido atualizar outro usuário"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var user models.Users

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	if err = repo.Update(user, ID); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// DeleteUser remove permanentemente um usuário do banco de dados pelo ID informado na URL.
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if userIDToken != ID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é permitido atualizar outro usuário"))
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	if err = repo.Delete(ID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

// FollowUser permite que um usuário siga outro usuário na aplicação.
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userID == followerID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é permitido que um usuário siga a si mesmo."))
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	if err = repo.Follow(userID, followerID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

// UnfollowUser permite que um usuário deixe de seguir outro usuário na aplicação.
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userID == followerID {
		responses.Erro(w, http.StatusInternalServerError, errors.New("Não é permitido que um usuário siga ou deixa de seguir a si mesmo."))
		return
	}

	db, err := data.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsersRepositories(db)
	if err = repo.Unfollow(userID, followerID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

// FindFollowers retorna todos os seguidores de um usuarios
func FindFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repo := repositories.NewUsersRepositories(db)
	followers, err := repo.FindFollowers(userID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// FindFollowing retorna os usuarios que estao sendo seguido pelo usuario que esta utilizado a apliacacao
func FindFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repo := repositories.NewUsersRepositories(db)
	following, err := repo.FindFollowing(userID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}
