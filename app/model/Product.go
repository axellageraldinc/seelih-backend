package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	TenantID            uint
	CategoryID          uint
	Sku                 string
	Name                string
	Quantity            uint
	PricePerItemPerDay  uint
	Description         string
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
	ProductStatus       string
	ImageName           string
}

const IMAGE_URL_PREFIX = "http://localhost:8080/api/products/img/"
