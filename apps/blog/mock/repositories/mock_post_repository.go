package repositories

import (
	"restful_blog/models"
)

type MockPostRepository struct {
	GetAllPostsFn  func(page, perPage int) ([]*models.Post, error)
	GetPostsByIDFn func(postsId int64) (*models.Post, error)
	CreatePostsFn  func(post *models.Post, categoryIDS []int64) error
	UpdatePostFn   func(post *models.Post, categoryIDs []int64) error
	DeletePostFn   func(postId int64) error
}

// GetAllPosts Mock get all posts
func (m *MockPostRepository) GetAllPosts(page, perPage int) ([]*models.Post, error) {
	return m.GetAllPostsFn(page, perPage)
}

// GetPostByID Mock get post by id
func (m *MockPostRepository) GetPostByID(postId int64) (*models.Post, error) {
	return m.GetPostsByIDFn(postId)
}

// CreatePost Mock create post
func (m *MockPostRepository) CreatePost(post *models.Post, categoryIDS []int64) error {
	if m.CreatePostsFn != nil {
		return m.CreatePostsFn(post, categoryIDS)
	}
	return nil
}

// UpdatePost Mock update post
func (m *MockPostRepository) UpdatePost(post *models.Post, categoryIDS []int64) error {
	return m.UpdatePostFn(post, categoryIDS)
}

// DeletePost Mock delete post
func (m *MockPostRepository) DeletePost(postId int64) error {
	return m.DeletePostFn(postId)
}
