package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"restful_blog/apps/blog/mock/repositories"
	"restful_blog/models"
	"restful_blog/service"
)

// Test_Category_Service_Get_Category_By_Id test category service GetCategoryById method
func Test_Category_Service_Get_Category_By_Id(t *testing.T) {
	// Create a mock repository
	mockRepo := &repositories.MockCategoryRepository{}

	// Set up the expected response from the mock repository
	expectedCategory := &models.Category{
		ID:   1,
		Name: "Category 1",
	}

	// Mock category repository
	mockRepo.GetCategoryByIDFn = func(categoryID int64) (*models.Category, error) {
		return expectedCategory, nil
	}

	// Create a new CategoryService instance with the mock  category repository
	s := &service.CategoryService{
		CategoryRepository: mockRepo,
	}

	// Call the method
	category, err := s.GetCategoryByID(expectedCategory.ID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

// Test_Category_Service_Get_All_Category test category service GetAllCategory
func Test_Category_Service_Get_All_Category(t *testing.T) {
	// Create a mock repository
	mockRepo := &repositories.MockCategoryRepository{}

	// Set up the expected response from the mock category repository
	expectedCategories := []*models.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}
	mockRepo.GetAllCategoriesFn = func(page, perPage int) ([]*models.Category, error) {
		return expectedCategories, nil
	}

	// Create a new CategoryService instance with the mock category repository
	s := &service.CategoryService{
		CategoryRepository: mockRepo,
	}

	// Call the method
	categories, err := s.GetAllCategory("1", "10")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
}

// Test_Category_Service_Create_Category test  category service CreateCategory service
func Test_Category_Service_Create_Category(t *testing.T) {
	// Create a mock category repository
	mockRepo := &repositories.MockCategoryRepository{}

	// Set up the expected response from the mock category repository
	mockRepo.CreateCategoryFn = func(category *models.Category) error {
		category.ID = 1 // Assign the expected ID
		return nil
	}

	// Create a new CategoryService instance with the mock category repository
	s := &service.CategoryService{
		CategoryRepository: mockRepo,
	}

	// Create a new category
	category := &models.Category{
		Name: "New Category",
	}

	// Call the method
	err := s.CreateCategory(category)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), category.ID)
}

// Test_Category_Service_Update_Category test category service UpdateCategory
func Test_Category_Service_Update_Category(t *testing.T) {
	// Create a mock repository
	mockRepo := &repositories.MockCategoryRepository{}

	// Set up the expected response from the mock category repository
	mockRepo.UpdateCategoryFn = func(category *models.Category) error {
		return nil
	}
	// Create a new CategoryService instance with the mock category repository
	s := &service.CategoryService{
		CategoryRepository: mockRepo,
	}
	// Create a category to update
	category := &models.Category{
		ID:   1,
		Name: "Updated Category",
	}

	// Call the method
	err := s.UpdateCategory(category)

	// Assertions
	assert.NoError(t, err)
}

// Test_Category_Service_Delete_Category test category service Delete method
func Test_Category_Service_Delete_Category(t *testing.T) {
	// Create a mock repository
	mockRepo := &repositories.MockCategoryRepository{}

	// Set up the expected response from the mock category repository
	mockRepo.DeleteCategoryFn = func(categoryID int64) error {
		return nil
	}

	// Create a new CategoryService instance with the mock category repository
	s := &service.CategoryService{
		CategoryRepository: mockRepo,
	}

	// Call the method
	err := s.DeleteCategory(1)

	// Assertions
	assert.NoError(t, err)
}
