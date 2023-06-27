package repositories

import (
	"restful_blog/models"
)

type MockCategoryRepository struct {
	GetAllCategoriesFn func(page, perPage int) ([]*models.Category, error)
	GetCategoryByIDFn  func(categoryId int64) (*models.Category, error)
	CreateCategoryFn   func(category *models.Category) error
	UpdateCategoryFn   func(category *models.Category) error
	DeleteCategoryFn   func(categoryId int64) error
}

// GetAllCategories Mock get all categories
func (m *MockCategoryRepository) GetAllCategories(page, perPage int) ([]*models.Category, error) {
	return m.GetAllCategoriesFn(page, perPage)
}

// GetCategoryByID Mock get category by id
func (m *MockCategoryRepository) GetCategoryByID(categoryId int64) (*models.Category, error) {
	return m.GetCategoryByIDFn(categoryId)
}

// CreateCategory Mock create category
func (m *MockCategoryRepository) CreateCategory(category *models.Category) error {
	if m.CreateCategoryFn != nil {
		return m.CreateCategoryFn(category)
	}
	return nil
}

// UpdateCategory Mock update category
func (m *MockCategoryRepository) UpdateCategory(category *models.Category) error {
	return m.UpdateCategoryFn(category)
}

// DeleteCategory Mock delete category
func (m *MockCategoryRepository) DeleteCategory(categoryId int64) error {
	return m.DeleteCategoryFn(categoryId)
}
