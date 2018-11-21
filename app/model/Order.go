package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ProductID uint
	BorrowerID uint
	Quantity uint
	TotalPrice uint
	OrderStatus string
	DeliveryType string
	RentDurationInDay uint
	ReturnTime time.Time
}
