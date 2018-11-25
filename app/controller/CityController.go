package controller

import (
	. "../helper"
	"encoding/json"
	"net/http"

	. "../mapper"
	. "../model/response"
	. "../service"
	"github.com/kataras/golog"
)

type CityController struct {
	ICityService
	ICityResponseMapper
}

func (cityController *CityController) GetAllCitiesHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/cities")
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(OK(cityController.ToCityResponseList(cityController.FindAll())))
}
