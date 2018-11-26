package service

import (
	"../mocks"
	"../model"
	"../util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderService_GetAllOrders_Success(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("FindAllOrdersByBorrowerId", 1).Return(util.OrdersResponse, 0)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := util.OrdersResponse
	expectedErrorCode := 0

	actualResult, actualErrorCode := orderService.GetAllOrders(1)

	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, uint(expectedErrorCode), actualErrorCode)
	assert.Equal(t, expectedResult[0].ProductName, actualResult[0].ProductName)
}

func TestOrderService_ConfirmProductRetrieval_Success(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.Order1.ID)).Return(true)
	orderRepository.On("FindOrderById", int(util.Order1.ID)).Return(util.Order1)
	orderRepository.On("UpdateOrderStatus", util.Order1, model.RETRIEVED)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := 0

	actualResult := orderService.ConfirmProductRetrieval(util.ConfirmProductRetrievalRequest1)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductRetrieval_Failed_OrderIdNotFound(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductRetrievalRequest2_OrderIdNotFound.OrderId)).Return(false)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_NOT_FOUND

	actualResult := orderService.ConfirmProductRetrieval(util.ConfirmProductRetrievalRequest2_OrderIdNotFound)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductReturn_Success(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductReturnRequest1.OrderId)).Return(true)
	orderRepository.On("FindOrderById", int(util.ConfirmProductReturnRequest1.OrderId)).Return(util.Order1)
	orderRepository.On("UpdateOrderStatus", util.Order1, model.DONE)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := 0

	actualResult := orderService.ConfirmProductReturn(util.ConfirmProductReturnRequest1)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductReturn_Failed_OrderIdNotFound(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductReturnRequest2_OrderIdNotFound.OrderId)).Return(false)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_NOT_FOUND

	actualResult := orderService.ConfirmProductReturn(util.ConfirmProductReturnRequest2_OrderIdNotFound)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductCancellation_Failed_OrderIdNotFound(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductCancellationRequest2_OrderIdNotFound.OrderId)).Return(false)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_NOT_FOUND

	actualResult := orderService.ConfirmProductCancellation(util.ConfirmProductCancellationRequest2_OrderIdNotFound)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductCancellation_Failed_OrderNotOnProcess(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductCancellationRequest1.OrderId)).Return(true)
	orderRepository.On("FindOrderById", int(util.ConfirmProductCancellationRequest1.OrderId)).Return(util.Order2_StatusNotOnProcess)
	productRepository.On("FindProductById", int(util.Order2_StatusNotOnProcess.ProductID)).Return(util.Product1)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_CANCELLATION_FAILED

	actualResult := orderService.ConfirmProductCancellation(util.ConfirmProductCancellationRequest1)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_ConfirmProductCancellation_Success(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	orderRepository.On("DoesOrderIdExist", int(util.ConfirmProductCancellationRequest1.OrderId)).Return(true)
	orderRepository.On("FindOrderById", int(util.ConfirmProductCancellationRequest1.OrderId)).Return(util.Order1)
	productRepository.On("FindProductById", int(util.Order1.ProductID)).Return(util.Product1)
	orderRepository.On("UpdateOrderStatus", util.Order1, model.CANCELLED)
	productRepository.On("UpdateProductStatus", util.Product1, model.OPENED)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := 0

	actualResult := orderService.ConfirmProductCancellation(util.ConfirmProductCancellationRequest1)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_OrderIdNotFound(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest1_OrderIdNotFound.ProductId)).Return(false)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_PRODUCT_ID_NOT_FOUND

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest1_OrderIdNotFound)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_ProductNotAvailableForRenting(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest2_OrderNotAvailableForRenting.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest2_OrderNotAvailableForRenting.ProductId)).Return(util.Product3_Closed)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_PRODUCT_NOT_AVAILABLE_FOR_RENTING

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest2_OrderNotAvailableForRenting)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_BorrowerIsTenant(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest3_BorrowerIsTenant.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest3_BorrowerIsTenant.ProductId)).Return(util.Product1)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_BORROWER_IS_THE_TENANT

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest3_BorrowerIsTenant)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_DurationDoesntMeetMinimum(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest4_DurationDoesntMeetMinimum.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest4_DurationDoesntMeetMinimum.ProductId)).Return(util.Product1)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_RENT_DURATION_DOESNT_MEET_MINIMUM_RENT_DURATION

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest4_DurationDoesntMeetMinimum)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_DurationExceedsMaximum(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest5_DurationExceedsMaximum.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest5_DurationExceedsMaximum.ProductId)).Return(util.Product1)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_RENT_DURATION_EXCEEDS_PRODUCT_MAX_RENT_DURATION

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest5_DurationExceedsMaximum)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_RequestedQuantityExceedsCurrentQuantity(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest6_RequestedQuantityExceedsCurrentQuantity.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest6_RequestedQuantityExceedsCurrentQuantity.ProductId)).Return(util.Product1)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_QUANTITY_EXCEEDS_PRODUCT_QUANTITY

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest6_RequestedQuantityExceedsCurrentQuantity)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Failed_BorrowerIdNotFound(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest7_BorrowerIdNotFound.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest7_BorrowerIdNotFound.ProductId)).Return(util.Product1)
	userRepository.On("DoesUserIdExist", int(util.PlaceOrderRequest7_BorrowerIdNotFound.BorrowerId)).Return(false)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := model.ORDER_FAILED_BORROWER_ID_NOT_FOUND

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest7_BorrowerIdNotFound)

	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestOrderService_PlaceOrder_Success_RemainingQuantityIsNotZero(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest8.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest8.ProductId)).Return(util.Product2)
	userRepository.On("DoesUserIdExist", int(util.PlaceOrderRequest8.BorrowerId)).Return(true)
	productRepository.On("UpdateProductQuantity", util.Product2, util.Product2.Quantity - util.PlaceOrderRequest8.Quantity)
	orderRepository.On("SaveOrder", mock.AnythingOfType("Order")).Return(true)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := 0

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest8)

	assert.Equal(t, expectedResult, int(actualResult))
}

func TestOrderService_PlaceOrder_Success_RemainingQuantityIsZero(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)

	productRepository.On("DoesProductIdExist", int(util.PlaceOrderRequest8.ProductId)).Return(true)
	productRepository.On("FindProductById", int(util.PlaceOrderRequest8.ProductId)).Return(util.Product1)
	userRepository.On("DoesUserIdExist", int(util.PlaceOrderRequest8.BorrowerId)).Return(true)
	productRepository.On("UpdateProductQuantityAndProductStatus", util.Product1, util.Product1.Quantity - util.PlaceOrderRequest8.Quantity, model.CLOSED)
	orderRepository.On("SaveOrder", mock.AnythingOfType("Order")).Return(true)

	orderService := OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}

	expectedResult := 0

	actualResult := orderService.PlaceOrder(util.PlaceOrderRequest8)

	assert.Equal(t, expectedResult, int(actualResult))
}