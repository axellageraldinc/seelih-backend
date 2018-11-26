package mocks

import (
	. "../model"
	. "../model/response"
	"github.com/stretchr/testify/mock"
)

type IOrderRepository struct {
	mock.Mock
}

func (mock *IOrderRepository) FindAllOrdersByBorrowerId(id int) (orders []OrderResponse, errorCode uint) {
	args := mock.Called(id)
	return args.Get(0).([]OrderResponse), uint(args.Int(1))
}

func (mock *IOrderRepository) DoesOrderIdExist(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}

func (mock *IOrderRepository) FindOrderById(id int) Order {
	args := mock.Called(id)
	return args.Get(0).(Order)
}

func (mock *IOrderRepository) UpdateOrderStatus(order Order, orderStatus string) {
	_ = mock.Called(order, orderStatus)
}

func (mock *IOrderRepository) SaveOrder(order Order) bool {
	args := mock.Called(order)
	return args.Bool(0)
}