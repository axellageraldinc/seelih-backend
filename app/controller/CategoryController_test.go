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

func TestCategoryController_FindAllCategoriesHandler_Success(t *testing.T) {
	// calling the mocked object
	categoryService := new(mocks.ICategoryService)
	categoryResponseMapper := new(mocks.ICategoryResponseMapper)

	// stubbing
	categoryService.On("FindAll").Return(util.Categories)
	categoryResponseMapper.On("ToCategoryResponseList", util.Categories).Return(util.CategoriesResponse)

	categoryController := CategoryController{
		ICategoryService:        categoryService,
		ICategoryResponseMapper: categoryResponseMapper,
	}

	// calling the tested method
	req := httptest.NewRequest("GET", "http://localhost:8080/api/categories", nil)
	w := httptest.NewRecorder()
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()
	routes.HandleFunc("/categories", categoryController.FindAllCategoriesHandler)
	routes.ServeHTTP(w, req)

	expectedResponse := response.WebResponse{
		HttpCode: 200,
		ErrorCode: 0,
		Data: util.CategoriesResponse,
	}

	var actualResponse response.WebResponse
	json.NewDecoder(w.Body).Decode(&actualResponse)

	// asserting
	assert.Equal(t, expectedResponse.HttpCode, actualResponse.HttpCode)
	assert.Equal(t, expectedResponse.ErrorCode, actualResponse.ErrorCode)
	assert.NotNil(t, actualResponse.Data)
	assert.NotEmpty(t, actualResponse.Data)
	var expectedData = expectedResponse.Data.([]response.CategoryResponse)
	var actualData = actualResponse.Data.([]interface{})
	assert.Equal(t, float64(expectedData[0].Id) , actualData[0].(map[string]interface{})["Id"])
	assert.Equal(t, float64(expectedData[1].Id) , actualData[1].(map[string]interface{})["Id"])
}
