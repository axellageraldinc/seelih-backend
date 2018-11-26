package mocks

import (
	. "../model/request"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IOrderService struct {
	mock.Mock
}

func (mock *IOrderService) PlaceOrder(request PlaceOrderRequest) uint {
	args := mock.Called(request)
	return uint(args.Int(0))
}

func (mock *IOrderService) GetAllOrders(userId int) (orders []OrderResponse, errorCode uint) {
	args := mock.Called(userId)
	return args.Get(0).([]OrderResponse), uint(args.Int(1))
}

func (mock *IOrderService) ConfirmProductRetrieval(request ConfirmProductRetrievalRequest) uint {
	args := mock.Called(request)
	return uint(args.Int(0))
}

func (mock *IOrderService) ConfirmProductReturn(request ConfirmProductReturnRequest) uint {
	args := mock.Called(request)
	return uint(args.Int(0))
}

func (mock *IOrderService) ConfirmProductCancellation(request ConfirmProductCancellationRequest) uint {
	args := mock.Called(request)
	return uint(args.Int(0))
}