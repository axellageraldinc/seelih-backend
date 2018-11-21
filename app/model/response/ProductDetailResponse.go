package response

import (
	"time"
)

type ProductDetailResponse struct {
	Id                  uint
	TenantId            uint
	CategoryId          uint
	Sku                 string
	Name                string
	Quantity            uint
	PricePerItemPerDay  uint
	Description         string
	UploadedTime        time.Time
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
	ImageUrl            string
}
