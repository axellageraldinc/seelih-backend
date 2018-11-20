package model

import (
	"github.com/jinzhu/gorm"
	"math/big"
)

type Product struct {
	gorm.Model
	TenantID uint
	CategoryID uint
	Sku string
	Name string
	Quantity uint
	PricePerItemPerDay big.Int
	Description string
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
	ProductStatus string
}