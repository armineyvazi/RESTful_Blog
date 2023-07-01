package model

import (
	"time"
)

type Category struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"not null"`
	CreatedAt *time.Time `json:"create_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime:true"`
}
