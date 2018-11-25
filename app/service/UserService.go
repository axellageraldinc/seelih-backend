package service

import (
	. "../model"
	. "../model/request"
	. "../repository"
	"github.com/kataras/golog"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	Register(RegisterRequest) (bool, uint)
	Login(LoginRequest) (User, uint)
	GetUserData(int) (User, uint)
}

type UserService struct {
	IUserRepository
}

func (userService *UserService) Register(request RegisterRequest) (bool, uint) {
	var user User
	if userService.DoesUserEmailExist(request.Email) {
		golog.Warn("User with email " + request.Email + " already exists in DB")
		return false, REGISTER_FAILED_EMAIL_ALREADY_EXISTS
	} else {
		plainPasswordInByte := ConvertPlainPasswordToByte(request.Password)
		hashedPassword := HashAndSalt(plainPasswordInByte)
		user = User{
			Email:       request.Email,
			Password:    hashedPassword,
			CityCodeId:  request.CityCode,
			Fullname:    request.Fullname,
			Fulladdress: request.FullAddress,
			Phone:       request.Phone,
		}
		if userService.SaveUser(user) {
			golog.Info("User registration succeed")
			return true, 0
		} else {
			golog.Warn("User registration failed")
			return false, REGISTER_FAILED_WONT_SAVE_TO_DATABASE
		}
	}
}

func (userService *UserService) Login(request LoginRequest) (User, uint) {
	if userService.DoesUserEmailExist(request.Email) {
		plainPassword := request.Password
		var userInDatabase = userService.FindUserByEmail(request.Email)
		userHashedPasswordFromDatabase := userInDatabase.Password
		isPasswordTrue := comparePasswords(userHashedPasswordFromDatabase, ConvertPlainPasswordToByte(plainPassword))
		if isPasswordTrue {
			golog.Info("Login succeed")
			return userInDatabase, 0
		} else {
			golog.Warn("Login failed, wrong password")
			return User{}, LOGIN_FAILED
		}
	} else {
		golog.Warn("Login failed, user with email " + request.Email + " not found")
		return User{}, LOGIN_FAILED
	}
}

func (userService *UserService) GetUserData(userId int) (User, uint) {
	if userService.DoesUserIdExist(userId) {
		var user = userService.FindUserById(userId)
		golog.Info("Get User Data Succeed")
		return user, 0
	} else {
		golog.Warn("No such user with id: " + string(userId))
		return User{}, LOGIN_FAILED
	}
}

func HashAndSalt(bytePlainPassword []byte) string {
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

func ConvertPlainPasswordToByte(plainPassword string) []byte {
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
