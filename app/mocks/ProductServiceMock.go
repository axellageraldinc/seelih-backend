package mocks

import (
	. "../model"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

type IProductService struct {
	mock.Mock
}

func (mock *IProductService) UploadProduct(productData string, file multipart.File, err error) uint {
	args := mock.Called(productData, file, err)
	return uint(args.Int(0))
}

func (mock *IProductService) GetAllAvailableProducts() []Product {
	args := mock.Called()
	return args.Get(0).([]Product)
}

func (mock *IProductService) GetOneProductDetails(id int) (product Product, errorCode uint) {
	args := mock.Called(id)
	return args.Get(0).(Product), uint(args.Int(1))
}

func (mock *IProductService) GetUserUploadedProducts(tenantId int) (products []Product, errorCode uint) {
	args := mock.Called(tenantId)
	return args.Get(0).([]Product), uint(args.Int(1))
}