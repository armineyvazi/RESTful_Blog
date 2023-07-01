package service

import (
	"strconv"

	"restful_blog/pkg/category/model"
	"restful_blog/pkg/category/repository"
)

type CategoryService struct {
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewCategoryService(cr repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		CategoryRepository: cr,
	}
}

func (c *CategoryService) CategoryByID(categoryId int64) (*model.Category, error) {
	return c.CategoryRepository.CategoryByID(categoryId)
}

func (c *CategoryService) AllCategory(pageStr, perPageStr string) ([]*model.Category, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	return c.CategoryRepository.AllCategories(page, perPage)

}

func (c *CategoryService) CreateCategory(category *model.Category) error {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryService) DeleteCategory(categoryID int64) error {
	return c.CategoryRepository.DeleteCategory(categoryID)
}
