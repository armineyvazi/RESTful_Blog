package service

import (
	"strconv"

	"restful_blog/models"
	"restful_blog/repositories"
)

type CategoryServiceInterface interface {
	GetCategoryByID(categoryId int64) (*models.Category, error)
	GetAllCategory(pageStr, perPageStr string) ([]*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID int64) error
}

type CategoryService struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewCategoryService() *CategoryService {
	categoryRepository := repositories.NewCategoryRepository()
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryService) GetCategoryByID(categoryId int64) (*models.Category, error) {
	return c.CategoryRepository.GetCategoryByID(categoryId)
}

func (c *CategoryService) GetAllCategory(pageStr, perPageStr string) ([]*models.Category, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	return c.CategoryRepository.GetAllCategories(page, perPage)

}

func (c *CategoryService) CreateCategory(category *models.Category) error {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryService) UpdateCategory(category *models.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryService) DeleteCategory(categoryID int64) error {
	return c.CategoryRepository.DeleteCategory(categoryID)
}
