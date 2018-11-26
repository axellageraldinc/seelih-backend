package service

import (
	. "../helper"
	. "../model"
	. "../model/request"
	. "../repository"
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

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
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE
	}
	productDataByte := []byte(productData)
	json.Unmarshal(productDataByte, &request)

	if request.PricePerItemPerDay <= 0 {
		return UPLOAD_PRODUCT_FAILED_PRICE_IS_ZERO_OR_BELOW
	} else if request.Quantity <= 0 {
		return UPLOAD_PRODUCT_FAILED_QUANTITY_IS_ZERO_OR_BELOW
	} else if !productService.DoesUserIdExist(int(request.TenantId)) {
		return UPLOAD_PRODUCT_FAILED_TENANT_ID_NOT_EXISTS
	} else if !productService.DoesCategoryIdExist(int(request.CategoryId)) {
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
	return uint(errorCode)
}

func uploadImage(file multipart.File, err error) (errorCode int, fileNameAndExtension string) {
	var defaultErrorCode = 0
	if err != nil {
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, ""
	}
	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" &&
		fileType != "image/jpg" &&
		fileType != "image/png" {
		return UPLOAD_PRODUCT_FAILED_INVALID_FILE, ""
	}
	fileName := uuid.Must(uuid.NewV4()).String()
	fileEndings, _ := mime.ExtensionsByType(fileType)
	newPath := filepath.Join(UPLOAD_PATH, fileName+fileEndings[0])
	newFile, _ := os.Create(newPath)
	defer newFile.Close()
	_, _ = newFile.Write(fileBytes)
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
		return []Product{}, GET_USER_UPLOADED_PRODUCTS_FAILED_USER_ID_NOT_FOUND
	} else {
		var products = productService.FindAllByTenantId(tenantId)
		return products, 0
	}
}