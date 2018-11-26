package service

import (
	. "../model"
	. "../model/request"
	. "../model/response"
	. "../repository"
	"time"
)

type IOrderService interface {
	PlaceOrder(request PlaceOrderRequest) uint
	GetAllOrders(userId int) (orders []OrderResponse, errorCode uint)
	ConfirmProductRetrieval(request ConfirmProductRetrievalRequest) uint
	ConfirmProductReturn(request ConfirmProductReturnRequest) uint
	ConfirmProductCancellation(request ConfirmProductCancellationRequest) uint
}

type OrderService struct {
	IOrderRepository
	IProductRepository
	IUserRepository
}

func (orderService *OrderService) PlaceOrder(request PlaceOrderRequest) uint {
	if !orderService.DoesProductIdExist(int(request.ProductId)) {
		return ORDER_FAILED_PRODUCT_ID_NOT_FOUND
	} else {
		product := orderService.FindProductById(int(request.ProductId))
		if product.ProductStatus == CLOSED {
			return ORDER_FAILED_PRODUCT_NOT_AVAILABLE_FOR_RENTING
		} else if request.BorrowerId == product.TenantID {
			return ORDER_FAILED_BORROWER_IS_THE_TENANT
		} else if request.Duration < product.MinimumBorrowedTime {
			return ORDER_FAILED_RENT_DURATION_DOESNT_MEET_MINIMUM_RENT_DURATION
		} else if request.Duration > product.MaximumBorrowedTime {
			return ORDER_FAILED_RENT_DURATION_EXCEEDS_PRODUCT_MAX_RENT_DURATION
		} else if request.Quantity > product.Quantity {
			return ORDER_FAILED_QUANTITY_EXCEEDS_PRODUCT_QUANTITY
		} else if !orderService.DoesUserIdExist(int(request.BorrowerId)) {
			return ORDER_FAILED_BORROWER_ID_NOT_FOUND
		} else {
			totalPrice := (product.PricePerItemPerDay * request.Quantity) * request.Duration
			remainingProductQuantity := product.Quantity - request.Quantity

			if remainingProductQuantity != 0 {
				orderService.UpdateProductQuantity(product, remainingProductQuantity)
			} else {
				orderService.UpdateProductQuantityAndProductStatus(product, remainingProductQuantity, CLOSED)
			}

			order := Order{
				ProductID:         request.ProductId,
				Quantity:          request.Quantity,
				BorrowerID:        request.BorrowerId,
				DeliveryType:      request.DeliveryType,
				OrderStatus:       ON_PROCESS,
				RentDurationInDay: request.Duration,
				TotalPrice:        totalPrice,
				ReturnTime:        time.Now().Add(time.Hour * 24 * time.Duration(request.Duration)),
			}
			orderService.SaveOrder(order)
			return 0
		}
	}
}

func (orderService *OrderService) GetAllOrders(userId int) (orders []OrderResponse, errorCode uint) {
	return orderService.FindAllOrdersByBorrowerId(userId)
}

func (orderService *OrderService) ConfirmProductRetrieval(request ConfirmProductRetrievalRequest) uint {
	if !orderService.DoesOrderIdExist(int(request.OrderId)) {
		return ORDER_NOT_FOUND
	} else {
		var order = orderService.FindOrderById(int(request.OrderId))
		orderService.UpdateOrderStatus(order, RETRIEVED)
		return 0
	}
}

func (orderService *OrderService) ConfirmProductReturn(request ConfirmProductReturnRequest) uint {
	if !orderService.DoesOrderIdExist(int(request.OrderId)) {
		return ORDER_NOT_FOUND
	} else {
		var order = orderService.FindOrderById(int(request.OrderId))
		orderService.UpdateOrderStatus(order, DONE)
		return 0
	}
}

func (orderService *OrderService) ConfirmProductCancellation(request ConfirmProductCancellationRequest) uint {
	if !orderService.DoesOrderIdExist(int(request.OrderId)) {
		return ORDER_NOT_FOUND
	} else {
		var order = orderService.FindOrderById(int(request.OrderId))
		var product = orderService.FindProductById(int(order.ProductID))
		if order.OrderStatus == ON_PROCESS {
			orderService.UpdateOrderStatus(order, CANCELLED)
			orderService.UpdateProductStatus(product, OPENED)
			return 0
		} else {
			return ORDER_CANCELLATION_FAILED
		}
	}
}