package repository

import (
	cm "restful_blog/pkg/category/model"
)

type CategoryRepositoryInterface interface {
	AllCategories(page, perPage int) ([]*cm.Category, error)
	CategoryByID(categoryId int64) (*cm.Category, error)
	CreateCategory(category *cm.Category) (*cm.Category, error)
	UpdateCategory(category *cm.Category) (*cm.Category, error)
	DeleteCategory(categoryId int64) error
}
