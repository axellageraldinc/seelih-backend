package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string
	Password string
	Fullname string
	Phone string
	City string
	Fulladdress string
	Products []Product `gorm:"foreignkey:TenantID"`
}