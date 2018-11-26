package util

import (
	"../model/response"
)

var UserResponse1 = response.UserResponse{
	Id: User1.ID,
	Phone: User1.Phone,
	Fullname: User1.Fullname,
	Fulladdress: User1.Fulladdress,
	CityCodeId: User1.CityCodeId,
	Email: User1.Email,
}