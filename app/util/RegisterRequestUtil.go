package util

import (
	"../model/request"
)

var RegisterRequest1 = request.RegisterRequest{
	Email: User1.Email,
	Password: "axell123",
	Phone: User1.Phone,
	Fullname: User1.Fullname,
	FullAddress: User1.Fulladdress,
	CityCode: User1.CityCodeId,
}