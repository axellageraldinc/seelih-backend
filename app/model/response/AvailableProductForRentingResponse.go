package response

import (
	"math/big"
	"time"
)

type AvailableProductForRentingResponse struct {
	Id                 uint
	Name               string
	PricePerItemPerDay big.Int
	UploadedTime       time.Time
}