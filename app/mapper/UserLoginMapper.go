package mapper

import (
	. "../model"
	. "../model/response"
)

type IUserLoginMapper interface {
	ToUserLoginDetail(User) UserResponse
}

type UserLoginMapper struct {}

func (userLoginMapper *UserLoginMapper) ToUserLoginDetail(user User) UserResponse {
	return UserResponse{
		Id:          user.ID,
		Fullname:    user.Fullname,
		Email:       user.Email,
		Phone:       user.Phone,
		Fulladdress: user.Fulladdress,
		CityCodeId:  user.CityCodeId,
	}
}
