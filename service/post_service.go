package service

import (
	"strconv"

	"restful_blog/models"
	"restful_blog/repositories"
)

type PostServiceInterface interface {
	GetPostByID(postID int64) (*models.Post, error)
	GetAllPosts(pageStr, perPageStr string) ([]*models.Post, error)
	CreatePost(post *models.Post, categoryIDs []int64) error
	UpdatePost(post *models.Post, categoryIDs []int64) error
	DeletePost(id int64) error
}

type PostService struct {
	PostRepository repositories.PostRepository
}

func NewPostService() *PostService {
	postRepository := repositories.NewPostRepository()
	return &PostService{
		PostRepository: postRepository,
	}
}

func (s *PostService) GetPostByID(postID int64) (*models.Post, error) {
	return s.PostRepository.GetPostByID(postID)
}

func (s *PostService) GetAllPosts(pageStr, perPageStr string) ([]*models.Post, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}
	return s.PostRepository.GetAllPosts(page, perPage)
}

func (s *PostService) CreatePost(post *models.Post, categoryIDs []int64) error {
	return s.PostRepository.CreatePost(post, categoryIDs)
}

func (s *PostService) UpdatePost(post *models.Post, categoryIDs []int64) error {
	return s.PostRepository.UpdatePost(post, categoryIDs)
}

func (s *PostService) DeletePost(id int64) error {
	return s.PostRepository.DeletePost(id)
}
