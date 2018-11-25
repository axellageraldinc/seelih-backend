package service

import (
	. "../model"
	. "../model/request"
	. "../model/response"
	. "../repository"
	"github.com/kataras/golog"
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
	var errorCode uint
	if !orderService.DoesProductIdExist(int(request.ProductId)) {
		errorCode = ORDER_FAILED_PRODUCT_ID_NOT_FOUND
	} else {
		product := orderService.FindProductById(int(request.ProductId))
		if product.ProductStatus == CLOSED {
			golog.Warn("Product is not available for renting")
			errorCode = ORDER_FAILED_PRODUCT_NOT_AVAILABLE_FOR_RENTING
		} else if request.BorrowerId == product.TenantID {
			golog.Warn("The borrower is the tenant, can't proceed!")
			errorCode = ORDER_FAILED_BORROWER_IS_THE_TENANT
		} else if request.Duration < product.MinimumBorrowedTime {
			golog.Warn("Rent duration requested by user doesn't meet the product's min rent duration")
			errorCode = ORDER_FAILED_RENT_DURATION_DOESNT_MEET_MINIMUM_RENT_DURATION
		} else if request.Duration > product.MaximumBorrowedTime {
			golog.Warn("Rent duration requested by user exceeds the product's max rent duration")
			errorCode = ORDER_FAILED_RENT_DURATION_EXCEEDS_PRODUCT_MAX_RENT_DURATION
		} else if request.Quantity > product.Quantity {
			golog.Warn("Quantity requested by user exceeds the product's quantity")
			errorCode = ORDER_FAILED_QUANTITY_EXCEEDS_PRODUCT_QUANTITY
		} else if !orderService.DoesUserIdExist(int(request.BorrowerId)) {
			golog.Warn("Borrower not exists in database")
			errorCode = ORDER_FAILED_BORROWER_ID_NOT_FOUND
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
			golog.Info("Order is placed successfully!")
			errorCode = 0
		}
	}
	return errorCode
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
		golog.Info("Confirm product retrieval succeed")
		return 0
	}
}

func (orderService *OrderService) ConfirmProductReturn(request ConfirmProductReturnRequest) uint {
	if !orderService.DoesOrderIdExist(int(request.OrderId)) {
		return ORDER_NOT_FOUND
	} else {
		var order = orderService.FindOrderById(int(request.OrderId))
		orderService.UpdateOrderStatus(order, DONE)
		golog.Info("Confirm product return succeed")
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
			order = orderService.FindOrderById(int(request.OrderId))
			orderService.UpdateOrderStatus(order, CANCELLED)
			orderService.UpdateProductStatus(product, OPENED)
			golog.Info("Confirm product cancellation succeed")
			return 0
		} else {
			golog.Warn("Order cancellation failed. Order is not ON_PROCESS.")
			return ORDER_CANCELLATION_FAILED
		}
	}
}