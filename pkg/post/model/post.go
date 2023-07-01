package model

import (
	"time"

	cm "restful_blog/pkg/category/model"
)

type Post struct {
	ID          int64          `json:"id" gorm:"primary_key"`
	Title       string         `json:"title" gorm:"not null" validate:"required"`
	Text        string         `json:"text" gorm:"not null"  validate:"required"`
	CategoryIDs []int64        `gorm:"-" json:"-" validate:"required"`
	CreatedAt   *time.Time     `json:"create_at" gorm:"autoCreateTime:true"`
	ModifiedAt  time.Time      `json:"modified_at" gorm:"autoUpdateTime:true"`
	Categories  []*cm.Category `gorm:"many2many:category_posts;"`
}
