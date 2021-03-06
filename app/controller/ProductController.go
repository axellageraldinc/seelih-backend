package controller

import (
	. "../helper"
	. "../mapper"
	. "../model"
	. "../model/response"
	. "../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ProductController struct {
	IProductService
	IAvailableProductForRentingResponseMapper
	IProductDetailResponseMapper
	IUploadedProductResponseMapper
}

func (productController *ProductController) UploadProductHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products POST")

	w.Header().Set(CONTENT_TYPE, MULTIPART_FORM_DATA)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)

	var response WebResponse

	r.Body = http.MaxBytesReader(w, r.Body, MAX_FILE_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_FILE_UPLOAD_SIZE); err != nil {
		response = ERROR(UPLOAD_PRODUCT_FAILED_FILE_TOO_BIG)
		json.NewEncoder(w).Encode(response)
		return
	}

	productData := r.FormValue("product_data")
	image, _, err := r.FormFile("image")

	errorCode := productController.UploadProduct(productData, image, err)
	if errorCode == 0 {
		response = OK(nil)
	} else {
		response = ERROR(errorCode)
	}

	json.NewEncoder(w).Encode(response)
}

func (productController *ProductController) GetAllAvailableProductsHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products GET")

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(OK(productController.ToAvailableProductForRentingResponseList(productController.GetAllAvailableProducts())))
}

func (productController *ProductController) GetOneProductDetailsHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products/{productId}")

	parameters := mux.Vars(r)
	productId, _ := strconv.ParseInt(parameters["productId"], 10, 32)

	var response WebResponse

	product, errorCode := productController.GetOneProductDetails(int(productId))
	if errorCode == 0 {
		response = OK(productController.ToProductDetailResponse(product))
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (productController *ProductController) GetUserUploadedProductsHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/users/{userId}/products")

	parameters := mux.Vars(r)
	userId, _ := strconv.ParseInt(parameters["userId"], 10, 32)

	var products []Product
	var response WebResponse

	products, errorCode := productController.GetUserUploadedProducts(int(userId))
	if errorCode == 0 {
		response = OK(productController.ToUploadedProductResponseList(products))
	} else {
		response = ERROR(errorCode)
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Header().Set(ACCESS_CONTROL_ALLOW_ORIGIN, ALL)
	json.NewEncoder(w).Encode(response)
}

func (productController *ProductController) GetProductImageHandler(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	imageName := parameters["imageName"]

	img, err := os.Open(UPLOAD_PATH + "/" + imageName)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set(CONTENT_TYPE, IMAGE_JPEG) // <-- set the content-type header
	io.Copy(w, img)
}
