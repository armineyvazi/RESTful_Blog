package repository

import (
	"gorm.io/gorm"

	"restful_blog/pkg/category/model"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (c *CategoryRepositoryImpl) CategoryByID(id int64) (*model.Category, error) {
	var category model.Category
	if err := c.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepositoryImpl) AllCategories(page, perPage int) ([]*model.Category, error) {
	var categories []*model.Category
	offset := (page - 1) * perPage
	err := c.db.Offset(offset).Limit(perPage).Model(categories).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryRepositoryImpl) CreateCategory(category *model.Category) (*model.Category, error) {
	if err := c.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) UpdateCategory(category *model.Category) (*model.Category, error) {
	if err := c.db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) DeleteCategory(categoryId int64) error {
	category := model.Category{ID: categoryId}

	if err := c.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
