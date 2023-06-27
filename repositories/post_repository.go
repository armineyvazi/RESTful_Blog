package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"restful_blog/database"
	"restful_blog/models"
)

type PostRepository interface {
	GetPostByID(postId int64) (*models.Post, error)
	GetAllPosts(page, perPage int) ([]*models.Post, error)
	CreatePost(post *models.Post, categoryIDs []int64) error
	UpdatePost(post *models.Post, categoryIDs []int64) error
	DeletePost(id int64) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository() *PostRepositoryImpl {
	db := database.GetDB()
	return &PostRepositoryImpl{
		db: db,
	}
}

func (r *PostRepositoryImpl) GetPostByID(postId int64) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, "id = ?", postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryImpl) GetAllPosts(page, perPage int) ([]*models.Post, error) {
	var posts []*models.Post
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Model(posts).Preload("Categories").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepositoryImpl) CreatePost(post *models.Post, categoryIDs []int64) error {
	var categories []*models.Category

	if err := r.db.Find(&categories, categoryIDs).Error; err != nil {
		return err
	}

	// Validate all category IDs exist.
	if len(categories) != len(categoryIDs) {
		return fmt.Errorf("categoryIDs exist")
	}

	post.Categories = categories

	if err := r.db.Create(post).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImpl) UpdatePost(post *models.Post, categoryIDs []int64) error {

	var categories []*models.Category

	if err := r.db.Find(&categories, categoryIDs).Error; err != nil {
		return err
	}

	// Validate all category IDs exist.
	if len(categories) != len(categoryIDs) {
		return fmt.Errorf("categoryIDs exist")
	}

	// Update the categories
	post.Categories = categories

	// Save the updated post.
	if err := r.db.Save(post).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImpl) DeletePost(id int64) error {
	// Remove the associations between the post and categories
	if err := r.db.Exec("DELETE FROM category_posts WHERE post_id = ?", id).Error; err != nil {
		return err
	}
	post := &models.Post{ID: id}

	if err := r.db.Delete(post).Error; err != nil {
		return err
	}

	return nil
}
