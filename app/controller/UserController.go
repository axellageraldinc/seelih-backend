package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../helper"
	"../mapper"
	. "../model"
	. "../model/request"
	. "../model/response"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/register")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var registerRequest RegisterRequest
	var user User
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&registerRequest)

	if db.Where("email = ?", registerRequest.Email).Find(&user).RecordNotFound() {
		plainPasswordInByte := convertPlainPasswordToByte(registerRequest.Password)
		hashedPassword := hashAndSalt(plainPasswordInByte)
		user = User{
			Email:       registerRequest.Email,
			Password:    hashedPassword,
			CityCodeId:  registerRequest.CityCode,
			Fullname:    registerRequest.Fullname,
			Fulladdress: registerRequest.FullAddress,
			Phone:       registerRequest.Phone,
		}
		if db.NewRecord(user) {
			db.Create(&user)
			if !db.NewRecord(user) {
				response = OK(&user)
				golog.Info("User registration succeed")
			} else {
				response = ERROR(0)
				golog.Info("User registration failed")
			}
		}
	} else {
		golog.Warn("User with email " + registerRequest.Email + " already exists in DB")
		response = ERROR(REGISTER_FAILED_EMAIL_ALREADY_EXISTS)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/login")

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
			db.Where("email = ?", loginRequest.Email).Find(&user)
			userLogin := mapper.ToUserLoginDetail(user)
			golog.Info("Login succeed")
			response = OK(userLogin)
		} else {
			golog.Warn("Login failed, wrong password")
			response = ERROR(LOGIN_FAILED)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/{userId} GET")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	parameters := mux.Vars(r)
	userId := parameters["userId"]

	var user User
	var response WebResponse

	if db.Where("id = ?", userId).Find(&user).RecordNotFound() {
		golog.Warn("No such user with id: " + userId)
		response = ERROR(LOGIN_FAILED)
	} else {
		db.Where("id = ?", userId).Find(&user)
		userData := mapper.ToUserLoginDetail(user)
		golog.Info("Get User Data Succeed")
		response = OK(userData)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func hashAndSalt(bytePlainPassword []byte) string {
	// Use GenerateFromPassword to hash & salt bytePlainPassword
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(bytePlainPassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
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
