package controller

import (
	"../helper"
	"../mapper"
	"../model"
	"../model/response"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/categories")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var categories []model.Category

	db.Find(&categories)

	json.NewEncoder(w).Encode(response.OK(mapper.ToCategoryResponseList(categories)))
}