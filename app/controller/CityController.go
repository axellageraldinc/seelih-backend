package controller

import (
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(OK(cityController.ToCityResponseList(cityController.FindAll())))
}
