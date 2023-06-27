package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"restful_blog/models"
)

// Test_category_model
func Test_category_model(t *testing.T) {
	// Create a sample category
	category := &models.Category{
		ID:        1,
		Name:      "Test Category",
		CreatedAt: nil,
		UpdatedAt: time.Now(),
	}

	// Assertions
	assert.Equal(t, int64(1), category.ID)
	assert.Equal(t, "Test Category", category.Name)
	assert.NotNil(t, category.UpdatedAt)
	assert.Empty(t, category.Posts)
}
