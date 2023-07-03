package service

import (
	"errors"
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

func (c *CategoryService) CreateCategory(category *model.Category) (*model.Category, error) {
	if category == nil {
		return nil, errors.New("the category is empty")
	}

	if category.Name == "" {
		return nil, errors.New("the name category is empty")
	}
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryService) UpdateCategory(category *model.Category) (*model.Category, error) {
	if category == nil {
		return nil, errors.New("the category is empty")
	}

	updatedCategory, err := c.CategoryRepository.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

func (c *CategoryService) DeleteCategory(categoryID int64) error {
	err := c.CategoryRepository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}
