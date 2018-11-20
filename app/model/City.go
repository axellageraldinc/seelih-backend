package model

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Name string
	Users []User `gorm:"foreignkey:CityId"`
}