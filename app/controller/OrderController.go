package controller

import (
	"../helper"
	"../model"
	. "../model/request"
	. "../model/response"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
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
					Quantity: remainingProductQuantity,
					ProductStatus: model.CLOSED,
				})
			}

			order := model.Order{
				ProductID: placeOrderRequest.ProductId,
				Quantity: placeOrderRequest.Quantity,
				BorrowerID: placeOrderRequest.BorrowerId,
				DeliveryType: placeOrderRequest.DeliveryType,
				OrderStatus: model.ON_PROCESS,
				RentDurationInDay: placeOrderRequest.Duration,
				TotalPrice: totalPrice,
			}
			db.Create(&order)
			response = OK(nil)
			golog.Info("Order is placed successfully!")
		}
	}
	json.NewEncoder(w).Encode(response)
}