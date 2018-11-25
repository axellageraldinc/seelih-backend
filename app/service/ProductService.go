package service

import (
	. "../model"
	. "../model/request"
	. "../repository"
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const MAX_FILE_UPLOAD_SIZE = 2 * 1024 * 1024 // 2 MB
const UPLOAD_PATH = "./img"

type IProductService interface {
	UploadProduct(productData string, file multipart.File, err error) uint
	GetAllAvailableProducts() []Product
	GetOneProductDetails(id int) (product Product, errorCode uint)
	GetUserUploadedProducts(tenantId int) (products []Product, errorCode uint)
}

type ProductService struct {
	IProductRepository
	IUserRepository
	ICategoryRepository
}

func (productService *ProductService) UploadProduct(productData string, file multipart.File, err error) uint {
	var request UploadProductRequest

	if productData == "" {
		golog.Warn("Invalid file product_data")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE
	}
	productDataByte := []byte(productData)
	json.Unmarshal(productDataByte, &request)

	if request.PricePerItemPerDay <= 0 {
		golog.Warn("Price specified is 0 or below")
		return UPLOAD_PRODUCT_FAILED_PRICE_IS_ZERO_OR_BELOW
	} else if request.Quantity <= 0 {
		golog.Warn("Quantity specified is 0 or below")
		return UPLOAD_PRODUCT_FAILED_QUANTITY_IS_ZERO_OR_BELOW
	} else if !productService.DoesUserIdExist(int(request.TenantId)) {
		golog.Warn("Tenant ID doesn't exist")
		return UPLOAD_PRODUCT_FAILED_TENANT_ID_NOT_EXISTS
	} else if !productService.DoesCategoryIdExist(int(request.CategoryId)) {
		golog.Warn("Category ID doesn't exist")
		return UPLOAD_PRODUCT_FAILED_CATEGORY_ID_NOT_EXISTS
	}
	errorCode, fileNameAndExtension := uploadImage(file, err)
	product := Product{
		TenantID:            request.TenantId,
		CategoryID:          request.CategoryId,
		Quantity:            request.Quantity,
		Name:                request.Name,
		Sku:                 request.Sku,
		Description:         request.Description,
		PricePerItemPerDay:  request.PricePerItemPerDay,
		MinimumBorrowedTime: request.MinimumBorrowedTime,
		MaximumBorrowedTime: request.MaximumBorrowedTime,
		ProductStatus:       OPENED,
		ImageName:           fileNameAndExtension,
	}
	productService.SaveProduct(product)
	if errorCode == 0 {
		golog.Info("Upload product succeed")
	}
	return 0
}

func uploadImage(file multipart.File, err error) (errorCode int, fileNameAndExtension string) {
	var defaultErrorCode = 0
	if err != nil {
		golog.Warn("Invalid file image FormFile")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, ""
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warn("Invalid file image ReadAll")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, ""
	}
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" &&
		fileType != "image/jpg" &&
		fileType != "image/png" {
		golog.Warn("Invalid file type. Only accepts jpeg, jpg, and png")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, ""
	}
	fileName := uuid.Must(uuid.NewV4()).String()
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		golog.Warn("Can't read file type")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, fileName + fileEndings[0]
	}
	newPath := filepath.Join(UPLOAD_PATH, fileName+fileEndings[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		golog.Warn("Can't write file JOIN")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, fileName + fileEndings[0]
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		golog.Warn("Can't write file WRITE")
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, fileName + fileEndings[0]
	}
	return defaultErrorCode, fileName + fileEndings[0]
}

func (productService *ProductService) GetAllAvailableProducts() []Product {
	return productService.FindAllByProductStatus(OPENED)
}

func (productService *ProductService) GetOneProductDetails(id int) (product Product, errorCode uint) {
	if productService.DoesProductIdExist(id) {
		return productService.FindProductById(id), 0
	} else {
		return Product{}, PRODUCT_NOT_FOUND
	}
}

func (productService *ProductService) GetUserUploadedProducts(tenantId int) (products []Product, errorCode uint) {
	if !productService.DoesUserIdExist(tenantId) {
		golog.Warn("User with ID " + string(tenantId) + " not found")
		return []Product{}, GET_USER_UPLOADED_PRODUCTS_FAILED_USER_ID_NOT_FOUND
	} else {
		var products = productService.FindAllByTenantId(tenantId)
		golog.Info("Getting user's uploaded products succeed")
		return products, 0
	}
}