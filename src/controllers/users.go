package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.PrepareToCreate(); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	user.Pass = ""
	response.JSON(w, http.StatusCreated, user)
}

// UpdateUser atualiza um usuário no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != tokenUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a user other than yours"))
		return
	}

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}
	user.ID = userID

	if error = user.PrepareToUpdate(); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	rowsAffected, error := repository.Update(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if rowsAffected == 0 {
		response.Error(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// FindUser busta todos os utuários do banco
func FindUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	user, error := repository.FindByID(userID)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}

	if user.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// FindUsers busca um usuário no banco
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	users, error := repository.Find(nameOrNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}

	response.JSON(w, http.StatusOK, users)
}

// DeleteUser remove um usuário do bando
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, error := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserID != userID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a user other than yours"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	if error := repository.Delete(userID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Follow follow the user
func Follow(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	folloeUserID, error := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if folloeUserID == authenticatedUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to follow yourself"))
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	if error = repository.Follow(authenticatedUserID, folloeUserID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
