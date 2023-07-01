package service

import (
	"strconv"

	"restful_blog/pkg/post/model"
	"restful_blog/pkg/post/repository"
)

type PostService struct {
	PostRepository repository.PostRepositoryInterface
}

func NewPostService(rp repository.PostRepositoryInterface) *PostService {
	return &PostService{
		PostRepository: rp,
	}
}

func (s *PostService) PostByID(postID int64) (*model.Post, error) {
	return s.PostRepository.PostByID(postID)
}

func (s *PostService) AllPosts(pageStr, perPageStr string) ([]*model.Post, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}
	return s.PostRepository.AllPosts(page, perPage)
}

func (s *PostService) CreatePost(post *model.Post, categoryIDs []int64) error {
	return s.PostRepository.CreatePost(post, categoryIDs)
}

func (s *PostService) UpdatePost(post *model.Post, categoryIDs []int64) error {
	return s.PostRepository.UpdatePost(post, categoryIDs)
}

func (s *PostService) DeletePost(id int64) error {
	return s.PostRepository.DeletePost(id)
}
