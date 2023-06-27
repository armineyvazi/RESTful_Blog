package service

import (
	"restful_blog/models"
)

type MockCategoryService struct {
	GetCategoryByIDFunc func(categoryID int64) (*models.Category, error)
	GetAllCategoryFunc  func(pageStr, perPageStr string) ([]*models.Category, error)
	CreateCategoryFunc  func(category *models.Category) error
	UpdateCategoryFunc  func(category *models.Category) error
	DeleteCategoryFunc  func(categoryID int64) error
}

// GetCategoryByID Mock get category by id
func (m *MockCategoryService) GetCategoryByID(categoryID int64) (*models.Category, error) {
	if m.GetCategoryByIDFunc != nil {
		return m.GetCategoryByIDFunc(categoryID)
	}
	return nil, nil
}

// GetAllCategory Mock get categories
func (m *MockCategoryService) GetAllCategory(pageStr, perPageStr string) ([]*models.Category, error) {
	if m.GetAllCategoryFunc != nil {
		return m.GetAllCategoryFunc(pageStr, perPageStr)
	}
	return nil, nil
}

// CreateCategory Mock create category
func (m *MockCategoryService) CreateCategory(category *models.Category) error {
	if m.CreateCategoryFunc != nil {
		return m.CreateCategoryFunc(category)
	}
	return nil
}

// UpdateCategory Mock update category
func (m *MockCategoryService) UpdateCategory(category *models.Category) error {
	if m.UpdateCategoryFunc != nil {
		return m.UpdateCategoryFunc(category)
	}
	return nil
}

// DeleteCategory Mock delete category
func (m *MockCategoryService) DeleteCategory(categoryID int64) error {
	if m.DeleteCategoryFunc != nil {
		return m.DeleteCategoryFunc(categoryID)
	}
	return nil
}
