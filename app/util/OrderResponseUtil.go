package util

import (
	"../model"
	"../model/response"
)

var OrderResponse1 = response.OrderResponse{
	TotalPrice: Order1.TotalPrice,
	OrderStatus: Order1.OrderStatus,
	ReturnTime: Order1.ReturnTime.String(),
	Id: Order1.ID,
	ImageUrl: model.IMAGE_URL_PREFIX + Product1.ImageName,
	ProductName: Product1.Name,
}
var OrdersResponse = []response.OrderResponse{
	OrderResponse1,
}