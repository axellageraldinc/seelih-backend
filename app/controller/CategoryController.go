package controller

import (
	. "../mapper"
	. "../model/response"
	. "../service"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
)

type CategoryController struct {
	ICategoryService
	ICategoryResponseMapper
}

func (categoryController *CategoryController) FindAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/categories")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(OK(categoryController.ToCategoryResponseList(categoryController.FindAll())))
}