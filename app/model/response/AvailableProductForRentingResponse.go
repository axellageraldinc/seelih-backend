package response

import (
	"time"
)

type AvailableProductForRentingResponse struct {
	Id                 uint
	Name               string
	PricePerItemPerDay uint
	UploadedTime       time.Time
}