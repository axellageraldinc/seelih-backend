package controller

import (
	"encoding/json"
	"net/http"

	"../helper"
	"../mapper"
	"../model"
	"../model/response"
	"github.com/kataras/golog"
)

func GetAllCities(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/city GET")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var cities []model.City

	db.Find(&cities)

	json.NewEncoder(w).Encode(response.OK(mapper.ToCityResponseList(cities)))
}
