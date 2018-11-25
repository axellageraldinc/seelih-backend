package repository

import (
	. "../helper"
	. "../model"
)

type IProductRepository interface {
	SaveProduct(product Product) bool
	FindAllByProductStatus(productStatus string) []Product
	FindProductById(id int) Product
	DoesProductIdExist(id int) bool
	UpdateProductQuantity(product Product, remainingQuantity uint)
	UpdateProductQuantityAndProductStatus(product Product, remainingQuantity uint, productStatus string)
	UpdateProductStatus(product Product, productStatus string)
}

type ProductRepository struct {
	IDatabaseConnectionHelper
}

func (productRepository *ProductRepository) SaveProduct(product Product) bool {
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Create(&product)
	return !db.NewRecord(product) // if the `product` object is not a new record, it means that it's successfully saved to database
}

func (productRepository *ProductRepository) FindAllByProductStatus(productStatus string) []Product {
	var products []Product
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Where("product_status = ?", OPENED).Find(&products)
	return products
}

func (productRepository *ProductRepository) FindProductById(id int) Product {
	var product Product
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Where("id = ?", id).Find(&product)
	return product
}

func (productRepository *ProductRepository) DoesProductIdExist(id int) bool {
	var product Product
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	return !db.Where("id = ?", id).Find(&product).RecordNotFound()
}

func (productRepository *ProductRepository) UpdateProductQuantity(product Product, remainingQuantity uint) {
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Model(&product).Update("quantity", remainingQuantity)
}

func (productRepository *ProductRepository) UpdateProductQuantityAndProductStatus(product Product, remainingQuantity uint, productStatus string) {
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Model(&product).Updates(Product{
		Quantity:      remainingQuantity,
		ProductStatus: productStatus,
	})
}

func (productRepository *ProductRepository) UpdateProductStatus(product Product, productStatus string) {
	db := productRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Model(&product).Update("product_status", productStatus)
}