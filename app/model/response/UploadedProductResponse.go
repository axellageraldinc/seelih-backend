package response

import "time"

type UploadedProductResponse struct {
	Id                 uint
	Name               string
	PricePerItemPerDay uint
	UploadedTime       time.Time
	ImageUrl           string
}