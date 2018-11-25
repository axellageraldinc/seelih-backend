package repository

import (
	. "../helper"
	. "../model"
)

type IUserRepository interface {
	SaveUser(user User) bool
	FindUserByEmail(email string) User
	FindUserById(id int) User
	DoesUserEmailExist(email string) bool
	DoesUserIdExist(id int) bool
}

type UserRepository struct {
	IDatabaseConnectionHelper
}

func (userRepository *UserRepository) SaveUser(user User) bool {
	db := userRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Create(&user)
	return !db.NewRecord(user) // if the `user` object is not a new record, it means that it's successfully saved to database
}

func (userRepository *UserRepository) FindUserByEmail(email string) User {
	db := userRepository.OpenDatabaseConnection()
	defer db.Close()
	var user User
	db.Where("email = ?", email).Find(&user)
	return user
}

func (userRepository *UserRepository) FindUserById(id int) User {
	db := userRepository.OpenDatabaseConnection()
	defer db.Close()
	var user User
	db.Where("id = ?", id).Find(&user)
	return user
}

func (userRepository *UserRepository) DoesUserEmailExist(email string) bool {
	db := userRepository.OpenDatabaseConnection()
	defer db.Close()
	var user User
	return !db.Where("email = ?", email).Find(&user).RecordNotFound()
}

func (userRepository *UserRepository) DoesUserIdExist(id int) bool {
	db := userRepository.OpenDatabaseConnection()
	defer db.Close()
	var user User
	return !db.Where("id = ?", id).Find(&user).RecordNotFound()
}