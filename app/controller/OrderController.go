package controller

import (
	. "../helper"
	. "../model/request"
	. "../model/response"
	. "../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"net/http"
	"strconv"
)

type OrderController struct {
	IOrderService
}

func (orderController *OrderController) PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders")

	var placeOrderRequest PlaceOrderRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&placeOrderRequest)

	errorCode := orderController.PlaceOrder(placeOrderRequest)
	if errorCode == 0 {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (orderController *OrderController) GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/{userId}/orders")

	var response WebResponse

	parameters := mux.Vars(r)
	userId, _ := strconv.ParseInt(parameters["userId"], 10, 32)

	orders, errorCode := orderController.GetAllOrders(int(userId))
	if errorCode == 0 {
		response = OK(orders)
	} else {
		response = ERROR(errorCode)
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (orderController *OrderController) ConfirmProductRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/retrieve")

	var confirmProductRetrievalRequest ConfirmProductRetrievalRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductRetrievalRequest)

	errorCode := orderController.ConfirmProductRetrieval(confirmProductRetrievalRequest)
	if errorCode == 0 {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (orderController *OrderController) ConfirmProductReturnHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/return")

	var confirmProductReturnRequest ConfirmProductReturnRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductReturnRequest)

	errorCode := orderController.ConfirmProductReturn(confirmProductReturnRequest)
	if errorCode == 0 {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (orderController *OrderController) ConfirmProductCancellationHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/cancellation")

	var confirmProductCancellationRequest ConfirmProductCancellationRequest
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductCancellationRequest)

	errorCode := orderController.ConfirmProductCancellation(confirmProductCancellationRequest)
	if errorCode == 0 {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}
