package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_post
func Test_post(t *testing.T) {
	// Create a sample post
	post := &Post{
		ID:          1,
		Title:       "Test Post",
		Text:        "This is a test post",
		CategoryIDs: []int64{1, 2, 3},
		CreatedAt:   nil,
		ModifiedAt:  time.Now(),
	}

	// Assertions
	assert.Equal(t, int64(1), post.ID)
	assert.Equal(t, "Test Post", post.Title)
	assert.Equal(t, "This is a test post", post.Text)
	assert.Equal(t, []int64{1, 2, 3}, post.CategoryIDs)
	assert.NotNil(t, post.ModifiedAt)
	assert.Empty(t, post.Categories)
}
