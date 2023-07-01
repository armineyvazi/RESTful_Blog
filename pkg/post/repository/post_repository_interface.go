package repository

import (
	pm "restful_blog/pkg/post/model"
)

type PostRepositoryInterface interface {
	PostByID(postId int64) (*pm.Post, error)
	AllPosts(page, perPage int) ([]*pm.Post, error)
	CreatePost(post *pm.Post, categoryIDs []int64) error
	UpdatePost(post *pm.Post, categoryIDs []int64) error
	DeletePost(id int64) error
}
