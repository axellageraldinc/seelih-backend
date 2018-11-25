package service

import (
	. "../model"
	. "../repository"
)

type ICategoryService interface {
	FindAll() []Category
}

type CategoryService struct {
	ICategoryRepository
}

func (categoryService *CategoryService) FindAll() []Category {
	return categoryService.FindAllCategories()
}