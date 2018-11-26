package controller

import (
	"../mocks"
	"../model/response"
	"../util"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCityController_GetAllCitiesHandler_Success(t *testing.T) {
	cityService := new(mocks.ICityService)
	cityResponseMapper := new(mocks.ICityResponseMapper)

	cityService.On("FindAll").Return(util.Cities)
	cityResponseMapper.On("ToCityResponseList", util.Cities).Return(util.CitiesResponse)

	cityController := CityController{
		ICityResponseMapper: cityResponseMapper,
		ICityService: cityService,
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/cities", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/cities", cityController.GetAllCitiesHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.CitiesResponse,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.([]response.CityResponse)
	var actualData = actualResponse.Data.([]interface{})
	assert.Equal(t, float64(expectedData[0].Id) , actualData[0].(map[string]interface{})["Id"])
	assert.Equal(t, float64(expectedData[1].Id) , actualData[1].(map[string]interface{})["Id"])
	assert.Equal(t, expectedData[0].Name , actualData[0].(map[string]interface{})["Name"])
	assert.Equal(t, expectedData[1].Name , actualData[1].(map[string]interface{})["Name"])
}
