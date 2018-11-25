package controller

import (
	. "../helper"
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
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(OK(categoryController.ToCategoryResponseList(categoryController.FindAll())))
}