package service

import (
	"../mocks"
	"../model"
	"../util"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestProductService_GetAllAvailableProducts_Success(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productRepository.On("FindAllByProductStatus", model.OPENED).Return(util.Products)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := util.Products

	actualResult := productService.GetAllAvailableProducts()

	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
}

func TestProductService_GetOneProductDetails_Success(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productRepository.On("DoesProductIdExist", 1).Return(true)
	productRepository.On("FindProductById", 1).Return(util.Product1)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := util.Product1
	expectedErrorCode := 0

	actualResult, actualErrorCode := productService.GetOneProductDetails(1)

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
	assert.Equal(t, uint(expectedErrorCode), actualErrorCode)
}

func TestProductService_GetOneProductDetails_Failed_ProductIdNoFound(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productRepository.On("DoesProductIdExist", 999).Return(false)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.Product{}
	expectedErrorCode := model.PRODUCT_NOT_FOUND

	actualResult, actualErrorCode := productService.GetOneProductDetails(999)

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
	assert.Equal(t, uint(expectedErrorCode), actualErrorCode)
}

func TestProductService_GetUserUploadedProducts_Success(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", 1).Return(true)
	productRepository.On("FindAllByTenantId", 1).Return(util.Products)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := util.Products
	expectedErrorCode := 0

	actualResult, actualErrorCode := productService.GetUserUploadedProducts(1)

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
	assert.Equal(t, uint(expectedErrorCode), actualErrorCode)
}

func TestProductService_GetUserUploadedProducts_Failed_UserIdNotFound(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", 999).Return(false)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedErrorCode := model.GET_USER_UPLOADED_PRODUCTS_FAILED_USER_ID_NOT_FOUND

	actualResult, actualErrorCode := productService.GetUserUploadedProducts(999)

	assert.NotNil(t, actualResult)
	assert.Empty(t, actualResult)
	assert.Equal(t, uint(expectedErrorCode), actualErrorCode)
}

func TestProductService_UploadProduct_Failed_InvalidProductData(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_INVALID_FILE

	actualResult := productService.UploadProduct("", nil, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestProductService_UploadProduct_Failed_PricePerItemPerDayIsZeroOrBelow(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_PRICE_IS_ZERO_OR_BELOW

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProductPricePerItemPerDayIsZeroOrBelow)

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), nil, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestProductService_UploadProduct_Failed_QuantityIsZeroOrBelow(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_QUANTITY_IS_ZERO_OR_BELOW

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProductPriceQuantityIsZeroOrBelow)

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), nil, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestProductService_UploadProduct_Failed_TenantIdNotFound(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", int(util.UploadProduct1.TenantId)).Return(false)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_TENANT_ID_NOT_EXISTS

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProduct1)

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), nil, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestProductService_UploadProduct_Failed_CategoryIdNotFound(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", int(util.UploadProduct1.TenantId)).Return(true)
	categoryRepository.On("DoesCategoryIdExist", int(util.UploadProduct1.CategoryId)).Return(false)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_CATEGORY_ID_NOT_EXISTS

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProduct1)

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), nil, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, uint(expectedResult), actualResult)
}

func TestProductService_UploadProduct_Success(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", int(util.UploadProduct1.TenantId)).Return(true)
	categoryRepository.On("DoesCategoryIdExist", int(util.UploadProduct1.CategoryId)).Return(true)
	productRepository.On("SaveProduct", mock.AnythingOfType("Product")).Return(true)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := 0

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProduct1)

	file, _ := os.Open("../../image_for_test.png")
	defer file.Close()

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), file, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult, int(actualResult))
}

func TestProductService_UploadProduct_Success_SaveImageFailed_FileNotImage(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", int(util.UploadProduct1.TenantId)).Return(true)
	categoryRepository.On("DoesCategoryIdExist", int(util.UploadProduct1.CategoryId)).Return(true)
	productRepository.On("SaveProduct", mock.AnythingOfType("Product")).Return(true)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_INVALID_FILE

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProduct1)

	file, _ := os.Create("test.txt")
	defer file.Close()

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), file, nil)

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult, int(actualResult))

	_ = os.Remove("test.txt")
}

func TestProductService_UploadProduct_Success_SaveImageFailed_ErrorExists(t *testing.T) {
	productRepository := new(mocks.IProductRepository)
	userRepository := new(mocks.IUserRepository)
	categoryRepository := new(mocks.ICategoryRepository)

	userRepository.On("DoesUserIdExist", int(util.UploadProduct1.TenantId)).Return(true)
	categoryRepository.On("DoesCategoryIdExist", int(util.UploadProduct1.CategoryId)).Return(true)
	productRepository.On("SaveProduct", mock.AnythingOfType("Product")).Return(true)

	productService := ProductService{
		ICategoryRepository: categoryRepository,
		IUserRepository: userRepository,
		IProductRepository: productRepository,
	}

	expectedResult := model.UPLOAD_PRODUCT_FAILED_INVALID_FILE

	jsonUploadProductRequest, _ := json.Marshal(util.UploadProduct1)

	file, _ := os.Open("../../image_for_test.png")
	defer file.Close()

	actualResult := productService.UploadProduct(string(jsonUploadProductRequest), file, errors.New("error"))

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult, int(actualResult))
}