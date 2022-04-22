package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Sign in user
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userSavedDatabase, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userSavedDatabase.Passwd, user.Passwd); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userSavedDatabase.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userSavedDatabase.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthData{ID: userID, Token: token})
}
