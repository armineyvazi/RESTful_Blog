package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"restful_blog/pkg/category/model"
)

type MockCategoryRepository struct {
	mock.Mock
}

// Mock CreateCategory
func (m *MockCategoryRepository) CreateCategory(category *model.Category) (*model.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Category), args.Error(1)
}

// Mock AllCategories
func (m *MockCategoryRepository) AllCategories(pageStr, perPageStr int) ([]*model.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]*model.Category), args.Error(1)
}

// Mock CategoryByID
func (m *MockCategoryRepository) CategoryByID(id int64) (*model.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Category), args.Error(1)
}

// Mock UpdateCategory
func (m *MockCategoryRepository) UpdateCategory(category *model.Category) (*model.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Category), args.Error(1)
}

// Mock DeleteCategory
func (m *MockCategoryRepository) DeleteCategory(categoryId int64) error {
	args := m.Called()
	return args.Error(0)
}

func Test_DeleteCategory(t *testing.T) {
	categoryID := int64(1)
	mockRepo := new(MockCategoryRepository)

	// Setup expectations
	mockRepo.On("DeleteCategory").Return(nil)
	testCategoryService := NewCategoryService(mockRepo)
	err := testCategoryService.DeleteCategory(categoryID)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Error Assertion
	assert.Nil(t, err)
}

func Test_UpdateCategory(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &model.Category{
		ID:   1,
		Name: "Updated Category",
	}

	// Setup expectations
	mockRepo.On("UpdateCategory").Return(category, nil)
	testCategoryService := NewCategoryService(mockRepo)
	updatedCategory, err := testCategoryService.UpdateCategory(category)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, category.ID, updatedCategory.ID)
	assert.Equal(t, category.Name, updatedCategory.Name)
	assert.Nil(t, err)
}

func Test_CreateCategory(t *testing.T) {
	var identifier int64 = 1
	mockRepo := new(MockCategoryRepository)

	category := &model.Category{
		ID:   identifier,
		Name: "test",
	}

	// Setup expectations
	mockRepo.On("CreateCategory").Return(category, nil)
	testCategoryService := NewCategoryService(mockRepo)
	result, err := testCategoryService.CreateCategory(category)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result.ID)
	assert.Equal(t, "test", result.Name)
	assert.Nil(t, err)
}

func Test_CategoryByID(t *testing.T) {
	var identifier int64 = 1
	mockRepo := new(MockCategoryRepository)

	category := &model.Category{
		ID:   identifier,
		Name: "test",
	}

	// Setup expectations
	mockRepo.On("CategoryByID").Return(category, nil)
	testCategoryService := NewCategoryService(mockRepo)
	result, err := testCategoryService.CategoryByID(identifier)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result.ID)
	assert.Equal(t, "test", result.Name)
	assert.Nil(t, err)
}

func Test_AllCategory(t *testing.T) {
	var identifier int64 = 1
	mockRepo := new(MockCategoryRepository)
	category := &model.Category{
		ID:   identifier,
		Name: "test",
	}
	// Setup expectations
	mockRepo.On("AllCategories").Return([]*model.Category{category}, nil)
	testCategoryService := NewCategoryService(mockRepo)
	result, err := testCategoryService.AllCategory("1", "2")

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "test", result[0].Name)
	assert.Nil(t, err)

}

func Test_CreateCategory_EmptyCategory(t *testing.T) {
	service := NewCategoryService(nil)
	data, err := service.CreateCategory(nil)

	assert.NotNil(t, err)
	assert.Nil(t, data)
	assert.Equal(t, "the category is empty", err.Error())
}

func Test_CreateCategory_CategoryName(t *testing.T) {
	service := NewCategoryService(nil)
	category := &model.Category{}
	data, err := service.CreateCategory(category)

	assert.NotNil(t, err)
	assert.Nil(t, data)
	assert.Equal(t, "the name category is empty", err.Error())

}
