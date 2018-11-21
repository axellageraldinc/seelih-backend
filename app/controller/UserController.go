package controller

import (
	"../helper"
	. "../model"
	. "../model/request"
	. "../model/response"
	"encoding/json"
	"github.com/kataras/golog"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/login")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var loginRequest LoginRequest
	var user User
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&loginRequest)

	if db.Where("email = ?", loginRequest.Email).Find(&user).RecordNotFound() {
		golog.Warn("Login failed, user with email " + loginRequest.Email + " not found")
		response = ERROR(LOGIN_FAILED)
	} else {
		plainPassword := loginRequest.Password
		db.Where("email = ?", loginRequest.Email).Find(&user)
		userHashedPasswordFromDatabase := user.Password
		isPasswordTrue := comparePasswords(userHashedPasswordFromDatabase, convertPlainPasswordToByte(plainPassword))
		if isPasswordTrue {
			golog.Info("Login succeed")
			response = OK(nil)
		} else {
			golog.Warn("Login failed, wrong password")
			response = ERROR(LOGIN_FAILED)
		}
	}
	json.NewEncoder(w).Encode(response)
}

func convertPlainPasswordToByte(plainPassword string) []byte {
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(plainPassword)
}

func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}