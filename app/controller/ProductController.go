package controller

import (
	"../helper"
	"../mapper"
	. "../model"
	. "../model/request"
	. "../model/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 MB
const UPLOAD_PATH = "./img"

func UploadProduct(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products POST")

	w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var uploadProductRequest UploadProductRequest
	//var user User
	//var category Category
	var response WebResponse

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		golog.Warn("File too big")
		response = ERROR(UPLOAD_PRODUCT_FAILED_FILE_TOO_BIG)
		json.NewEncoder(w).Encode(response)
		return
	}

	file, _, err := r.FormFile("product_data")
	if err != nil {
		golog.Warn("Invalid file product_data")
		response = ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	json.Unmarshal(fileBytes, &uploadProductRequest)

	//if uploadProductRequest.PricePerItemPerDay <= 0 {
	//	golog.Warn("Price specified is 0 or below")
	//	response = ERROR(UPLOAD_PRODUCT_FAILED_PRICE_IS_ZERO_OR_BELOW)
	//	json.NewEncoder(w).Encode(response)
	//	return
	//} else if uploadProductRequest.Quantity <= 0 {
	//	golog.Warn("Quantity specified is 0 or below")
	//	response = ERROR(UPLOAD_PRODUCT_FAILED_QUANTITY_IS_ZERO_OR_BELOW)
	//	json.NewEncoder(w).Encode(response)
	//	return
	//} else if db.Where("id = ?", uploadProductRequest.TenantId).Find(&user).RecordNotFound() {
	//	golog.Warn("Tenant ID doesn't exist")
	//	response = ERROR(UPLOAD_PRODUCT_FAILED_TENANT_ID_NOT_EXISTS)
	//	json.NewEncoder(w).Encode(response)
	//	return
	//} else if db.Where("id = ?", uploadProductRequest.CategoryId).Find(&category).RecordNotFound() {
	//	golog.Warn("Category ID doesn't exist")
	//	response = ERROR(UPLOAD_PRODUCT_FAILED_CATEGORY_ID_NOT_EXISTS)
	//	json.NewEncoder(w).Encode(response)
	//	return
	//}

	response, fileNameAndExtension := uploadImage(w, r)
	product := Product{
		TenantID:            uploadProductRequest.TenantId,
		CategoryID:          uploadProductRequest.CategoryId,
		Quantity:            uploadProductRequest.Quantity,
		Name:                uploadProductRequest.Name,
		Sku:                 uploadProductRequest.Sku,
		Description:         uploadProductRequest.Description,
		PricePerItemPerDay:  uploadProductRequest.PricePerItemPerDay,
		MinimumBorrowedTime: uploadProductRequest.MinimumBorrowedTime,
		MaximumBorrowedTime: uploadProductRequest.MaximumBorrowedTime,
		ProductStatus:       OPENED,
		ImageName:           fileNameAndExtension,
	}
	db.Create(&product)
	if response.ErrorCode == 0 {
		golog.Info("Upload product succeed")
	}
	json.NewEncoder(w).Encode(response)
}

func uploadImage(w http.ResponseWriter, r *http.Request) (webResponse WebResponse, fileNameAndExtension string) {
	var response = OK(nil)
	file, _, err := r.FormFile("image")
	if err != nil {
		golog.Warn("Invalid file image FormFile")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), ""
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warn("Invalid file image ReadAll")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), ""
	}
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" &&
		fileType != "image/jpg" &&
		fileType != "image/png" {
		golog.Warn("Invalid file type. Only accepts jpeg, jpg, and png")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), ""
	}
	fileName := uuid.Must(uuid.NewV4()).String()
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		golog.Warn("Can't read file type")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), fileName + fileEndings[0]
	}
	newPath := filepath.Join(UPLOAD_PATH, fileName+fileEndings[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		golog.Warn("Can't write file JOIN")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), fileName + fileEndings[0]
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		golog.Warn("Can't write file WRITE")
		return ERROR(UPLOAD_PRODUCT_FAILED_INVALID_FILE), fileName + fileEndings[0]
	}
	return response, fileName + fileEndings[0]
}

func GetAllAvailableProducts(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products GET")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	var products []Product

	db.Where("product_status = ?", OPENED).Find(&products)

	availableProductForRentingResponseList := mapper.ToAvailableProductForRentingResponseList(products)

	response := OK(availableProductForRentingResponseList)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetOneProductDetails(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/products/{productId}")

	db := helper.OpenDatabaseConnection()
	defer db.Close()

	parameters := mux.Vars(r)
	productId := parameters["productId"]

	var product Product
	var response WebResponse

	if db.Where("id = ?", productId).Find(&product).RecordNotFound() {
		golog.Warn("product with ID " + productId + " not found!")
		response = ERROR(PRODUCT_NOT_FOUND)
	} else {
		productDetailResponse := mapper.ToProductDetailResponse(product)
		response = OK(productDetailResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
