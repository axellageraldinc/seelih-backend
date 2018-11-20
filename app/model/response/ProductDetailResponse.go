package response

import (
	"math/big"
	"time"
)

type ProductDetailResponse struct {
	Id uint
	TenantId uint
	CategoryId uint
	Sku string
	Name string
	Quantity uint
	PricePerItemPerDay big.Int
	Description string
	UploadedTime time.Time
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
}
