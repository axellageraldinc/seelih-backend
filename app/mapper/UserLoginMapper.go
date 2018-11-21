package mapper

import (
	. "../model"
	. "../model/response"
)

func ToUserLoginDetail(user User) UserResponse {
	return UserResponse{
		Id:          user.ID,
		Fullname:    user.Fullname,
		Email:       user.Email,
		Phone:       user.Phone,
		Fulladdress: user.Fulladdress,
		CityCodeId:  user.CityCodeId,
	}
}
