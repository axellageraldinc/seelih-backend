package model

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Name string
	code uint
	Users []User `gorm:"foreignkey:CityId"`
}