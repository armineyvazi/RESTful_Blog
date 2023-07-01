package service

import (
	pm "restful_blog/pkg/post/model"
)

type PostServiceInterface interface {
	PostByID(postID int64) (*pm.Post, error)
	AllPosts(pageStr, perPageStr string) ([]*pm.Post, error)
	CreatePost(post *pm.Post, categoryIDs []int64) error
	UpdatePost(post *pm.Post, categoryIDs []int64) error
	DeletePost(id int64) error
}
