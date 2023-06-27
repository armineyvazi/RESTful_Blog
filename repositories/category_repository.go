package repositories

import (
	"gorm.io/gorm"

	"restful_blog/database"
	"restful_blog/models"
)

type CategoryRepositoryInterface interface {
	GetAllCategories(page, perPage int) ([]*models.Category, error)
	GetCategoryByID(categoryId int64) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryId int64) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	db := database.GetDB()
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (c *CategoryRepositoryImpl) GetCategoryByID(id int64) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepositoryImpl) GetAllCategories(page, perPage int) ([]*models.Category, error) {
	var categories []*models.Category
	offset := (page - 1) * perPage
	err := c.db.Offset(offset).Limit(perPage).Model(categories).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryRepositoryImpl) CreateCategory(category *models.Category) error {
	if err := c.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (c *CategoryRepositoryImpl) UpdateCategory(category *models.Category) error {
	if err := c.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (c *CategoryRepositoryImpl) DeleteCategory(categoryId int64) error {
	category := models.Category{ID: categoryId}

	if err := c.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
