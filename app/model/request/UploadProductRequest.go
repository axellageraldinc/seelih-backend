package request

type UploadProductRequest struct {
	TenantId uint
	CategoryId uint
	Sku string
	Name string
	Quantity uint
	PricePerItemPerDay uint
	Description string
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
}