package response

type OrderResponse struct {
	Id uint
	ProductName string
	ImageUrl string
	TotalPrice uint
	OrderStatus string
	ReturnTime string
}