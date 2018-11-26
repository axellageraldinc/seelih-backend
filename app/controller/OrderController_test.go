package controller

import (
	"../mocks"
	"../model"
	"../model/response"
	"../util"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestOrderController_GetAllOrdersHandler_Success(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("GetAllOrders", 1).Return(util.OrdersResponse, 0)

	orderController := OrderController{
		IOrderService: orderService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/1/orders", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}/orders", orderController.GetAllOrdersHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.OrdersResponse,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.([]response.OrderResponse)
	var actualData = actualResponse.Data.([]interface{})
	assert.Equal(t, float64(expectedData[0].Id) , actualData[0].(map[string]interface{})["Id"])
}

func TestOrderController_GetAllOrdersHandler_Failed_ErrorCodeExists(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("GetAllOrders", 999).Return([]response.OrderResponse{}, model.GET_ALL_ORDERS_FAILED_SQL_ERROR)

	orderController := OrderController{
		IOrderService: orderService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users/999/orders", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/users/{userId}/orders", orderController.GetAllOrdersHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.GET_ALL_ORDERS_FAILED_SQL_ERROR,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductRetrievalHandler_Success(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductRetrieval", util.ConfirmProductRetrievalRequest1).Return(0)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductRetrievalRequest, _ := json.Marshal(util.ConfirmProductRetrievalRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/retrieve", bytes.NewBuffer(jsonConfirmProductRetrievalRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/retrieve", orderController.ConfirmProductRetrievalHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductRetrievalHandler_Failed_ErrorCodeExists(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductRetrieval", util.ConfirmProductRetrievalRequest1).Return(model.ORDER_NOT_FOUND)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductRetrievalRequest, _ := json.Marshal(util.ConfirmProductRetrievalRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/retrieve", bytes.NewBuffer(jsonConfirmProductRetrievalRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/retrieve", orderController.ConfirmProductRetrievalHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.ORDER_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductReturnHandler_Success(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductReturn", util.ConfirmProductReturnRequest1).Return(0)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductReturnRequest, _ := json.Marshal(util.ConfirmProductReturnRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/return", bytes.NewBuffer(jsonConfirmProductReturnRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/return", orderController.ConfirmProductReturnHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductReturnHandler_Failed_ErrorCodeExists(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductReturn", util.ConfirmProductReturnRequest1).Return(model.ORDER_NOT_FOUND)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductReturnRequest, _ := json.Marshal(util.ConfirmProductReturnRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/return", bytes.NewBuffer(jsonConfirmProductReturnRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/return", orderController.ConfirmProductReturnHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.ORDER_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductCancellationHandler_Success(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductCancellation", util.ConfirmProductCancellationRequest1).Return(0)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductCancellationRequest, _ := json.Marshal(util.ConfirmProductCancellationRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/cancellation", bytes.NewBuffer(jsonConfirmProductCancellationRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/cancellation", orderController.ConfirmProductCancellationHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_ConfirmProductCancellationHandler_Failed_ErrorCodeExists(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("ConfirmProductCancellation", util.ConfirmProductCancellationRequest1).Return(model.ORDER_NOT_FOUND)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonConfirmProductCancellationRequest, _ := json.Marshal(util.ConfirmProductCancellationRequest1)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders/cancellation", bytes.NewBuffer(jsonConfirmProductCancellationRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders/cancellation", orderController.ConfirmProductCancellationHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.ORDER_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_PlaceOrderHandler_Success(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("PlaceOrder", util.PlaceOrderRequest8).Return(0)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonPlaceOrderRequest, _ := json.Marshal(util.PlaceOrderRequest8)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders", bytes.NewBuffer(jsonPlaceOrderRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders", orderController.PlaceOrderHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}

func TestOrderController_PlaceOrderHandler_Failed_ErrorCodeExists(t *testing.T) {
	orderService := new(mocks.IOrderService)

	orderService.On("PlaceOrder", util.PlaceOrderRequest8).Return(model.ORDER_FAILED_PRODUCT_ID_NOT_FOUND)

	orderController := OrderController{
		IOrderService: orderService,
	}

	jsonPlaceOrderRequest, _ := json.Marshal(util.PlaceOrderRequest8)

	req := httptest.NewRequest("POST", "http://localhost:8080/api/orders", bytes.NewBuffer(jsonPlaceOrderRequest))
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/orders", orderController.PlaceOrderHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: model.ORDER_FAILED_PRODUCT_ID_NOT_FOUND,
		Data: nil,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.Nil(t, actualResponse.Data)
}