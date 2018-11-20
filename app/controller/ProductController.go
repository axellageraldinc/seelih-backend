package controller

import (
	"../helper"
	"../mapper"
	. "../model"
	. "../model/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"net/http"
)

func GetAllAvailableProducts(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var products []Product

	db.Where("product_status = ?", OPENED).Find(&products)

	availableProductForRentingResponseList := mapper.ToAvailableProductForRentingResponseList(products)

	response := OK(availableProductForRentingResponseList)

	json.NewEncoder(w).Encode(response)
}

func GetOneProductDetails(w http.ResponseWriter, r *http.Request)  {
	golog.Info("/api/products/{productId}")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	parameters := mux.Vars(r)
	productId := parameters["productId"]

	var product Product
	var response WebResponse

	if db.Where("id = ?", productId).Find(&product).RecordNotFound() {
		golog.Warn("product with ID " + productId + " not found!")
		response = ERROR(PRODUCT_NOT_FOUND)
	} else {
		productDetailResponse := mapper.ToProductDetailResponse(product)
		response = OK(productDetailResponse)
	}

	json.NewEncoder(w).Encode(response)
}