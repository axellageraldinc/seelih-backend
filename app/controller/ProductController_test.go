package controller

import (
	"../mocks"
	"../model"
	"../model/response"
	"../util"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestProductController_GetAllAvailableProductsHandler_Success(t *testing.T) {
	productService := new(mocks.IProductService)
	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)

	productService.On("GetAllAvailableProducts").Return(util.Products)
	availableProductForRentingResponseMapper.On("ToAvailableProductForRentingResponseList", util.Products).Return(util.AvailableProductsForRentingResponse)

	productController := ProductController{
		IUploadedProductResponseMapper: uploadedProductResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductService: productService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/products", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/products", productController.GetAllAvailableProductsHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.AvailableProductsForRentingResponse,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.([]response.AvailableProductForRentingResponse)
	var actualData = actualResponse.Data.([]interface{})
	assert.Equal(t, float64(expectedData[0].Id) , actualData[0].(map[string]interface{})["Id"])
	assert.Equal(t, float64(expectedData[1].Id) , actualData[1].(map[string]interface{})["Id"])
}

func TestProductController_GetOneProductDetailsHandler_Success(t *testing.T) {
	productService := new(mocks.IProductService)
	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)

	productService.On("GetOneProductDetails", 1).Return(util.Product1, 0)
	productDetailResponseMapper.On("ToProductDetailResponse", util.Product1).Return(util.ProductDetailResponse1)

	productController := ProductController{
		IUploadedProductResponseMapper: uploadedProductResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductService: productService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/products/1", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/products/{productId}", productController.GetOneProductDetailsHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.ProductDetailResponse1,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.(response.ProductDetailResponse)
	var actualData = actualResponse.Data.(interface{})
	assert.Equal(t, float64(expectedData.Id) , actualData.(map[string]interface{})["Id"])
	assert.Equal(t, expectedData.Name , actualData.(map[string]interface{})["Name"])
}

func TestProductController_GetOneProductDetailsHandler_Failed_ErrorCodeExists(t *testing.T) {
	productService := new(mocks.IProductService)
	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)

	productService.On("GetOneProductDetails", 999).Return(util.Product1, model.PRODUCT_NOT_FOUND)

	productController := ProductController{
		IUploadedProductResponseMapper: uploadedProductResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductService: productService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/products/999", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/products/{productId}", productController.GetOneProductDetailsHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.PRODUCT_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestProductController_GetUserUploadedProductsHandler_Success(t *testing.T) {
	productService := new(mocks.IProductService)
	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)

	productService.On("GetUserUploadedProducts", 1).Return(util.Products, 0)
	uploadedProductResponseMapper.On("ToUploadedProductResponseList", util.Products).Return(util.UploadedProductsResponse)

	productController := ProductController{
		IUploadedProductResponseMapper: uploadedProductResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductService: productService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/1/products", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}/products", productController.GetUserUploadedProductsHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.UploadedProductsResponse,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.([]response.UploadedProductResponse)
	var actualData = actualResponse.Data.([]interface{})
	assert.Equal(t, float64(expectedData[0].Id) , actualData[0].(map[string]interface{})["Id"])
	assert.Equal(t, float64(expectedData[1].Id) , actualData[1].(map[string]interface{})["Id"])
}

func TestProductController_GetUserUploadedProductsHandler_Failed_ErrorCodeExists(t *testing.T) {
	productService := new(mocks.IProductService)
	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)

	productService.On("GetUserUploadedProducts", 999).Return([]model.Product{}, model.GET_USER_UPLOADED_PRODUCTS_FAILED_USER_ID_NOT_FOUND)

	productController := ProductController{
		IUploadedProductResponseMapper: uploadedProductResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductService: productService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/999/products", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}/products", productController.GetUserUploadedProductsHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.GET_USER_UPLOADED_PRODUCTS_FAILED_USER_ID_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

//func TestProductController_UploadProductHandler_Success(t *testing.T) {
//	productService := new(mocks.IProductService)
//	availableProductForRentingResponseMapper := new(mocks.IAvailableProductForRentingResponseMapper)
//	productDetailResponseMapper := new(mocks.IProductDetailResponseMapper)
//	uploadedProductResponseMapper := new(mocks.IUploadedProductResponseMapper)
//
//	productService.On("UploadProduct", mock.AnythingOfType("string"), mock.AnythingOfType("multipart.File"), mock.AnythingOfType("error")).Return(0)
//
//	productController := ProductController{
//		IUploadedProductResponseMapper: uploadedProductResponseMapper,
//		IProductDetailResponseMapper: productDetailResponseMapper,
//		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
//		IProductService: productService,
//	}
//
//	uploadProductRequest, _ := json.Marshal("uploadProductRequest")
//
//	req := httptest.NewRequest("POST", "http://localhost:8080/api/products", bytes.NewBuffer(uploadProductRequest))
//	req.Header.Set("Content-Type", "multipart/form-data")
//	w := httptest.NewRecorder()
//	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
//	routes.HandleFunc("/products", productController.UploadProductHandler)
//	routes.ServeHTTP(w, req)
//
//	expectedResponse := response.WebResponse{
//		HttpCode: 200,
//		ErrorCode: 0,
//		Data: nil,
//	}
//
//	var actualResponse response.WebResponse
//	json.NewDecoder(w.Body).Decode(&actualResponse)
//
//	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
//	assert.Equal(t, int(expectedResponse.ErrorCode), int(actualResponse.ErrorCode))
//}