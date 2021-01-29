package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login authenticate a user in API
func Login(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error := json.Unmarshal(body, &user); error != nil {
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

	persistedUser, error := repository.FindByEmail(user.Email)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error := security.VerifyPassword(persistedUser.Pass, user.Pass); error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, _ := authentication.CreateToken(persistedUser.ID)
	response.JSON(w, http.StatusOK, token)
}
