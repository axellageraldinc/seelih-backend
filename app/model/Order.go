package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ProductID uint
	BorrowerID uint
	OrderTime time.Time
	Quantity uint
	TotalPrice uint
	OrderStatus string
	DeliveryType string
}
