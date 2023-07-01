package repository

import (
	"fmt"

	"gorm.io/gorm"

	"restful_blog/pkg/category/model"
	pm "restful_blog/pkg/post/model"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (r *PostRepositoryImpl) PostByID(postId int64) (*pm.Post, error) {
	var post pm.Post
	if err := r.db.First(&post, "id = ?", postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryImpl) AllPosts(page, perPage int) ([]*pm.Post, error) {
	var posts []*pm.Post
	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Model(posts).Preload("Categories").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepositoryImpl) CreatePost(post *pm.Post, categoryIDs []int64) error {
	var categories []*model.Category

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

func (r *PostRepositoryImpl) UpdatePost(post *pm.Post, categoryIDs []int64) error {

	var categories []*model.Category

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
	post := &pm.Post{ID: id}

	if err := r.db.Delete(post).Error; err != nil {
		return err
	}

	return nil
}
