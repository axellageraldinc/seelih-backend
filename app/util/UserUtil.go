package util

import (
	"../model"
)

var User1 = model.User{
	Email: "email1",
	Password: "$2a$04$6Im3uNXceFfVlT74qQ3ezeoFDB8v8ju./LiXSc1/G1fjfPvKRvXVu", //axell123
	CityCodeId: 1,
	Fullname: "fullname1",
	Fulladdress: "fulladdress1",
	Phone: "phone1",
}
var User2 = model.User{
	Email: "email2",
	Password: "$2a$04$6Im3uNXceFfVlT74qQ3ezeoFDB8v8ju./LiXSc1/G1fjfPvKRvXVu", //axell123
	CityCodeId: 2,
	Fullname: "fullname2",
	Fulladdress: "fulladdress2",
	Phone: "phone2",
}
var NewUser = model.User{
	Email: RegisterRequest1.Email,
	CityCodeId: RegisterRequest1.CityCode,
	Password: "$2a$04$6Im3uNXceFfVlT74qQ3ezeoFDB8v8ju./LiXSc1/G1fjfPvKRvXVu", //axell123
	Fulladdress: RegisterRequest1.FullAddress,
	Fullname: RegisterRequest1.Fullname,
	Phone: RegisterRequest1.Phone,
}