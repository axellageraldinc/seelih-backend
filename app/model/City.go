package model

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Name string
	Code uint
	Users []User `gorm:"foreignkey:CityId"`
}