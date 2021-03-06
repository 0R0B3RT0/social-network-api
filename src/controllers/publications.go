package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

//CreatePublication create a new publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	authenticatedUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if publication.UserID != authenticatedUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to create a publication to other user than yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	if err = publication.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	publicationID, err := repository.Create(publication)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	publication, err = repository.Find(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

//UpdatePublication update an existing publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	publicationID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	authenticatedUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	publicationPersisted, err := repository.Find(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationPersisted.UserID != authenticatedUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to update a publication from other user"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	err = json.Unmarshal(body, &publication)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if _, err = repository.Update(publicationID, publication); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

//FindPublications find all publications from the user
func FindPublications(w http.ResponseWriter, r *http.Request) {
	authenticatedUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	publications, err := repository.FindByUserAndFollowUsers(authenticatedUserID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publications)
}

//FindPublication find a specific publication
func FindPublication(w http.ResponseWriter, r *http.Request) {
	publicationID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = authentication.ValidateToken(r); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	publication, err := repository.Find(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publication.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New("publication not found"))
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

//DeletePublication remove the publication from the database
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	pubID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	authenticatedUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	publicationPersisted, err := repository.Find(pubID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if publicationPersisted.UserID != authenticatedUserID {
		response.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a publication from other user"))
		return
	}

	if err = repository.Delete(pubID); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindPublicationByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	publications, err := repository.FindByUser(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	if _, err = repository.LikePublication(id); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnlikePublication(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepositories(db)

	if _, err = repository.UnlikePublication(id); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
