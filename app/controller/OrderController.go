package controller

import (
	"../helper"
	"../model"
	. "../model/request"
	. "../model/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"net/http"
	"time"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var placeOrderRequest PlaceOrderRequest
	var product model.Product
	var borrower model.User
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&placeOrderRequest)

	if db.Where("id = ?", placeOrderRequest.ProductId).Find(&product).RecordNotFound() {
		golog.Warn("Product with ID " + string(placeOrderRequest.ProductId) + " not found!")
		response = ERROR(model.ORDER_FAILED_PRODUCT_ID_NOT_FOUND)
	} else {
		db.Where("id = ?", placeOrderRequest.ProductId).Find(&product)
		if product.ProductStatus == model.CLOSED {
			golog.Warn("Product is not available for renting")
			response = ERROR(model.ORDER_FAILED_PRODUCT_NOT_AVAILABLE_FOR_RENTING)
		} else if placeOrderRequest.BorrowerId == product.TenantID {
			golog.Warn("The borrower is the tenant, can't proceed!")
			response = ERROR(model.ORDER_FAILED_BORROWER_IS_THE_TENANT)
		} else if placeOrderRequest.Duration < product.MinimumBorrowedTime {
			golog.Warn("Rent duration requested by user doesn't meet the product's min rent duration")
			response = ERROR(model.ORDER_FAILED_RENT_DURATION_DOESNT_MEET_MINIMUM_RENT_DURATION)
		} else if placeOrderRequest.Duration > product.MaximumBorrowedTime {
			golog.Warn("Rent duration requested by user exceeds the product's max rent duration")
			response = ERROR(model.ORDER_FAILED_RENT_DURATION_EXCEEDS_PRODUCT_MAX_RENT_DURATION)
		} else if placeOrderRequest.Quantity > product.Quantity {
			golog.Warn("Quantity requested by user exceeds the product's quantity")
			response = ERROR(model.ORDER_FAILED_QUANTITY_EXCEEDS_PRODUCT_QUANTITY)
		} else if db.Where("id = ?", placeOrderRequest.BorrowerId).Find(&borrower).RecordNotFound() {
			golog.Warn("Borrower not exists in database")
			response = ERROR(model.ORDER_FAILED_BORROWER_ID_NOT_FOUND)
		} else {
			totalPrice := (product.PricePerItemPerDay * placeOrderRequest.Quantity) * placeOrderRequest.Duration
			remainingProductQuantity := placeOrderRequest.Quantity - placeOrderRequest.Quantity

			if remainingProductQuantity != 0 {
				db.Model(&product).Update("quantity", remainingProductQuantity)
			} else {
				db.Model(&product).Updates(model.Product{
					Quantity:      remainingProductQuantity,
					ProductStatus: model.CLOSED,
				})
			}

			order := model.Order{
				ProductID:         placeOrderRequest.ProductId,
				Quantity:          placeOrderRequest.Quantity,
				BorrowerID:        placeOrderRequest.BorrowerId,
				DeliveryType:      placeOrderRequest.DeliveryType,
				OrderStatus:       model.ON_PROCESS,
				RentDurationInDay: placeOrderRequest.Duration,
				TotalPrice:        totalPrice,
				ReturnTime:        time.Now().Add(time.Hour * 24 * time.Duration(placeOrderRequest.Duration)),
			}
			db.Create(&order)
			response = OK(nil)
			golog.Info("Order is placed successfully!")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/{userId}/orders")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var user model.User
	var orderResponses []OrderResponse
	var response WebResponse

	parameters := mux.Vars(r)
	userId := parameters["userId"]

	if db.Where("id = ?", userId).Find(&user).RecordNotFound() {
		golog.Warn("User with ID " + userId + " not found")
		response = ERROR(model.GET_ALL_ORDERS_FAILED_USER_ID_NOT_FOUND)
	} else {
		rows, err := db.Raw("SELECT orders.id, products.name AS product_name, products.image_name AS image_url, orders.total_price, orders.order_status, orders.return_time "+
			"FROM orders, products "+
			"WHERE orders.borrower_id = ? AND orders.product_id = products.id", userId).Rows()
		defer rows.Close()
		if err != nil {
			golog.Warn("Error raw SQL selecting all orders " + err.Error())
			response = ERROR(model.GET_ALL_ORDERS_FAILED_SQL_ERROR)
		} else {
			for rows.Next() {
				var orderResponse OrderResponse
				db.ScanRows(rows, &orderResponse)
				orderResponse.ImageUrl = model.IMAGE_URL_PREFIX + orderResponse.ImageUrl
				orderResponses = append(orderResponses, orderResponse)
			}
			response = OK(orderResponses)
			golog.Info("Get all orders succeed")
		}
	}
	json.NewEncoder(w).Encode(response)
}

func ConfirmProductRetrieval(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/retrieve")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var confirmProductRetrievalRequest ConfirmProductRetrievalRequest
	var order model.Order
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductRetrievalRequest)

	if db.Where("id = ?", confirmProductRetrievalRequest.OrderId).Find(&order).RecordNotFound() {
		golog.Warn("Order with ID " + string(confirmProductRetrievalRequest.OrderId) + " not found!")
		response = ERROR(model.ORDER_NOT_FOUND)
	} else {
		db.Where("id = ?", confirmProductRetrievalRequest.OrderId).Find(&order)
		db.Model(&order).Update("order_status", model.RETRIEVED)
		response = OK(nil)
		golog.Info("Confirm product retrieval succeed")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func ConfirmProductReturn(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/return")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var confirmProductReturnRequest ConfirmProductReturnRequest
	var order model.Order
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductReturnRequest)

	if db.Where("id = ?", confirmProductReturnRequest.OrderId).Find(&order).RecordNotFound() {
		golog.Warn("Order with ID " + string(confirmProductReturnRequest.OrderId) + " not found!")
		response = ERROR(model.ORDER_NOT_FOUND)
	} else {
		db.Where("id = ?", confirmProductReturnRequest.OrderId).Find(&order)
		db.Model(&order).Update("order_status", model.DONE)
		response = OK(nil)
		golog.Info("Confirm product return succeed")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func ConfirmProductCancellation(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/orders/cancellation")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var confirmProductCancellationRequest ConfirmProductCancellationRequest
	var order model.Order
	var product model.Product
	var response WebResponse

	json.NewDecoder(r.Body).Decode(&confirmProductCancellationRequest)

	if db.Where("id = ?", confirmProductCancellationRequest.OrderId).Find(&order).RecordNotFound() {
		golog.Warn("Order with ID " + string(confirmProductCancellationRequest.OrderId) + " not found!")
		response = ERROR(model.ORDER_NOT_FOUND)
	} else {
		db.Where("id = ?", confirmProductCancellationRequest.OrderId).Find(&order)
		db.Where("id = ?", order.ProductID).Find(&product)
		if order.OrderStatus == model.ON_PROCESS {
			db.Where("id = ?", confirmProductCancellationRequest.OrderId).Find(&order)
			db.Model(&order).Update("order_status", model.CANCELLED)
			db.Model(&product).Update("product_status", model.OPENED)
			response = OK(nil)
			golog.Info("Confirm product cancellation succeed")
		} else {
			response = ERROR(model.ORDER_CANCELLATION_FAILED)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
