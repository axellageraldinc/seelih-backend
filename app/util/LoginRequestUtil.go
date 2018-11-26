package util

import (
	"../model/request"
)

var LoginRequest1 = request.LoginRequest{
	Email: User1.Email,
	Password: "axell123",
}
var LoginRequest2 = request.LoginRequest{
	Email: User1.Email,
	Password: "wrong password",
}