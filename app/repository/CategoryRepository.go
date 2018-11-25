package repository

import (
	. "../helper"
	. "../model"
)

type ICategoryRepository interface {
	FindAllCategories() []Category
	DoesCategoryIdExist(id int) bool
}

type CategoryRepository struct {
	IDatabaseConnectionHelper
}

func (categoryRepository *CategoryRepository) FindAllCategories() []Category {
	var categories []Category
	db := categoryRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Find(&categories)
	return categories
}

func (categoryRepository *CategoryRepository) DoesCategoryIdExist(id int) bool {
	var category Category
	db := categoryRepository.OpenDatabaseConnection()
	defer db.Close()
	return !db.Where("id = ?", id).Find(&category).RecordNotFound()
}