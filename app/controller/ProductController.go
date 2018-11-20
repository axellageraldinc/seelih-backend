package controller

import (
	"../helper"
	"../mapper"
	. "../model"
	. "../model/response"
	"encoding/json"
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