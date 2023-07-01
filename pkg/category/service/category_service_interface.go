package service

import (
	cm "restful_blog/pkg/category/model"
)

type CategoryServiceInterface interface {
	CategoryByID(categoryId int64) (*cm.Category, error)
	AllCategory(pageStr, perPageStr string) ([]*cm.Category, error)
	CreateCategory(category *cm.Category) error
	UpdateCategory(category *cm.Category) error
	DeleteCategory(categoryID int64) error
}
