package service

import (
	"../mocks"
	"../util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryService_FindAll_Success(t *testing.T) {
	categoryRepository := new(mocks.ICategoryRepository)

	categoryRepository.On("FindAllCategories").Return(util.Categories)

	categoryService := CategoryService{
		ICategoryRepository: categoryRepository,
	}

	expectedResult := util.Categories

	actualResult := categoryService.FindAll()

	assert.NotNil(t, actualResult)
	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
}
