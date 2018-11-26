package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
)

type IProductRepository struct {
	mock.Mock
}

func (mock *IProductRepository) SaveProduct(product Product) bool {
	args := mock.Called(product)
	return args.Bool(0)
}

func (mock *IProductRepository) FindAllByProductStatus(productStatus string) []Product {
	args := mock.Called(productStatus)
	return args.Get(0).([]Product)
}

func (mock *IProductRepository) FindAllByTenantId(tenantId int) []Product {
	args := mock.Called(tenantId)
	return args.Get(0).([]Product)
}

func (mock *IProductRepository) FindProductById(id int) Product {
	args := mock.Called(id)
	return args.Get(0).(Product)
}

func (mock *IProductRepository) DoesProductIdExist(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}

func (mock *IProductRepository) UpdateProductQuantity(product Product, remainingQuantity uint) {
	_ = mock.Called(product, remainingQuantity)
}

func (mock *IProductRepository) UpdateProductQuantityAndProductStatus(product Product, remainingQuantity uint, productStatus string) {
	_ = mock.Called(product, remainingQuantity, productStatus)
}

func (mock *IProductRepository) UpdateProductStatus(product Product, productStatus string) {
	_ = mock.Called(product, productStatus)
}