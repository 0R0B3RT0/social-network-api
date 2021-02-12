package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareToCreate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.Pass = ""
	response.JSON(w, http.StatusCreated, user)
}

// UpdateUser update user attributes
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

	followedUserID, error := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if followedUserID == authenticatedUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to follow yourself"))
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepositories(db)

	if error = repository.Follow(authenticatedUserID, followedUserID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Unfollow unfollow an user
func Unfollow(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, error := authentication.ExtractUserID(r)
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	followedUserID, error := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if followedUserID == authenticatedUserID {
		response.JSON(w, http.StatusNoContent, nil)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	repository := repositories.NewUserRepositories(db)

	if error = repository.Unfollow(authenticatedUserID, followedUserID); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func Followers(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadGateway, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users := repositories.NewUserRepositories(db)

	followers, err := users.Followers(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(followers) == 0 {
		response.Error(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func Followings(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadGateway, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users := repositories.NewUserRepositories(db)

	followings, err := users.Followings(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(followings) == 0 {
		response.Error(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, followings)
}

//PasswordUpdate update user password
func PasswordUpdate(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	if authenticatedUserID != userID {
		response.Error(w, http.StatusForbidden, errors.New("you can not update another user's password"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var pass models.Password
	if err = json.Unmarshal(body, &pass); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	hashNewPass, err := security.Hash(pass.NewPassword)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	pass.NewPassword = string(hashNewPass)

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users := repositories.NewUserRepositories(db)

	password, err := users.GetUserPassword(authenticatedUserID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(password, pass.LastPassword)
	if err != nil {
		log.Printf("Invalid password, userID: %d", userID)
		response.Error(w, http.StatusForbidden, errors.New("invalid user or password"))
	}

	isUpdated, err := users.UpdatePassword(authenticatedUserID, pass.NewPassword)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if !isUpdated {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
