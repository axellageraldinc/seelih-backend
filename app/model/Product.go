package model

type Product struct {
	TenantID uint
	CategoryID uint
	Sku string
	Name string
	Quantity uint
	PricePerItemPerDay uint
	Description string
	MinimumBorrowedTime uint
	MaximumBorrowedTime uint
}