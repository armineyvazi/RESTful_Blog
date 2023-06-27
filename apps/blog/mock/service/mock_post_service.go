package service

import (
	"fmt"

	"restful_blog/models"
)

type MockPostService struct {
	GetPostByIDFunc func(postId int64) (*models.Post, error)
	GetAllPostFunc  func(pageStr, perPageStr string) ([]*models.Post, error)
	CreatePostFunc  func(post *models.Post, categoryIDS []int64) error
	UpdatePostFunc  func(post *models.Post, categoryIDs []int64) error
	DeletePostFunc  func(postId int64) error
}

// GetPostByID Mock get post by id
func (m *MockPostService) GetPostByID(postId int64) (*models.Post, error) {
	return m.GetPostByIDFunc(postId)
}

// GetAllPosts Mock get all post
func (m *MockPostService) GetAllPosts(pageStr, perPageStr string) ([]*models.Post, error) {
	if m.GetAllPostFunc != nil {
		return m.GetAllPostFunc(pageStr, perPageStr)
	}
	return nil, nil
}

// CreatePost Mock create post
func (m *MockPostService) CreatePost(post *models.Post, categoryIDs []int64) error {
	return m.CreatePostFunc(post, categoryIDs)
}

// UpdatePost Mock update post
func (m *MockPostService) UpdatePost(post *models.Post, categoryIDs []int64) error {
	fmt.Println("armin")
	return m.UpdatePostFunc(post, categoryIDs)
}

// DeletePost Mock delete post
func (m *MockPostService) DeletePost(postId int64) error {
	return m.DeletePostFunc(postId)
}
