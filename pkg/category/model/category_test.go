package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_category_model
func Test_category_model(t *testing.T) {
	// Create a sample category
	category := &Category{
		ID:        1,
		Name:      "Test Category",
		CreatedAt: nil,
		UpdatedAt: time.Now(),
	}

	// Assertions
	assert.Equal(t, int64(1), category.ID)
	assert.Equal(t, "Test Category", category.Name)
	assert.NotNil(t, category.UpdatedAt)
}
