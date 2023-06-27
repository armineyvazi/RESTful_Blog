package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"restful_blog/models"
	"restful_blog/repositories"
)

// Test_Post_Repository_Get_Post_By_Id test post repository GetPostById method
func Test_Post_Repository_Get_Post_By_Id(t *testing.T) {
	var categories []int64

	repo := repositories.NewPostRepository()
	repoCategory := repositories.NewCategoryRepository()

	// Create a category to be retrieved
	expectedCategory := &models.Category{
		ID:   104,
		Name: "Category 1",
	}

	//Create categories id
	categories = append(categories, expectedCategory.ID)

	// insert category in database
	error := repoCategory.CreateCategory(expectedCategory)
	assert.NoError(t, error)

	// Create a category to be retrieved
	expectedPost := &models.Post{
		Title:       "Armin",
		Text:        "Armin",
		CategoryIDs: nil,
	}

	// insert post in database
	error = repo.CreatePost(expectedPost, categories)
	assert.NoError(t, error)

	// Call the method
	post, err := repo.GetPostByID(expectedPost.ID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedPost.ID, post.ID)

	// clean database
	err = repo.DeletePost(expectedPost.ID)
	assert.NoError(t, err)

	// clean database
	err = repoCategory.DeleteCategory(expectedCategory.ID)
	assert.NoError(t, err)

}

// Test_Post_Repository_Get_All_Post test post repository GetAllPost method
func Test_Post_Repository_Get_All_Post(t *testing.T) {
	// create array for save categoryIDs to array and pass to CreatePost method
	var categories []int64

	// Create a new CategoryRepositoryImpl instance
	postRepo := repositories.NewPostRepository()

	// Create some posts in the database
	expectedPost1 := &models.Post{
		ID:          105,
		Title:       "armin",
		Text:        "armin2",
		CategoryIDs: nil,
	}

	expectedPost2 := &models.Post{
		ID:          106,
		Title:       "armin",
		Text:        "armin2",
		CategoryIDs: nil,
	}

	// Create category
	repoCategory := repositories.NewCategoryRepository()

	// Create a category to be retrieved
	newCategory := &models.Category{
		ID:   104,
		Name: "Category 1",
	}

	//Create categories id
	categories = append(categories, newCategory.ID)

	//Create new category
	error := repoCategory.CreateCategory(newCategory)
	assert.NoError(t, error)

	// insert expected post 1 in database
	error = postRepo.CreatePost(expectedPost1, categories)
	assert.NoError(t, error)

	// insert expected post 2 in database
	error = postRepo.CreatePost(expectedPost2, categories)
	assert.NoError(t, error)

	// Call the method
	post, err := postRepo.GetAllPosts(1, 2)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedPost1.ID, post[0].ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost2.ID, post[1].ID)

	// Remove post in database
	err = postRepo.DeletePost(expectedPost1.ID)
	assert.NoError(t, err)
	err = postRepo.DeletePost(expectedPost2.ID)
	assert.NoError(t, err)

	// Remove category in database.
	err = repoCategory.DeleteCategory(newCategory.ID)
	assert.NoError(t, err)

}

// Test_Post_Repository_Create_Post test post repository CreatePost method
func Test_Post_Repository_Create_Post(t *testing.T) {
	// Create a new Post Repository instance
	repo := repositories.NewPostRepository()

	// Create some posts in the database
	post := &models.Post{
		ID:          105,
		Title:       "armin",
		Text:        "armin2",
		CategoryIDs: nil,
	}

	// insert post in database.
	err := repo.CreatePost(post, nil)

	// Assertions
	assert.NoError(t, err)
	assert.NotZero(t, post.ID)

	// clean database
	err = repo.DeletePost(post.ID)
	assert.NoError(t, err)
}

// Test_Post_Repository_UpdatePost test post repository UpdatePost method
func Test_Post_Repository_UpdatePost(t *testing.T) {
	// Create a new post Repository instance
	repo := repositories.NewPostRepository()

	// Create a post to update
	post := &models.Post{
		ID:          105,
		Title:       "armin",
		Text:        "armin2",
		CategoryIDs: nil,
	}

	error := repo.CreatePost(post, nil)
	assert.NoError(t, error)

	// Call the method
	err := repo.UpdatePost(post, nil)

	// Assertions
	assert.NoError(t, err)

	err = repo.DeletePost(post.ID)
	assert.NoError(t, err)
}

// Test_Post_Repository_DeletePost test post repository DeletePost method
func Test_Post_Repository_DeletePost(t *testing.T) {
	// Create a new post Repository instance
	repo := repositories.NewPostRepository()

	// Create a post to update
	post := &models.Post{
		ID:          105,
		Title:       "armin",
		Text:        "armin2",
		CategoryIDs: nil,
	}

	// insert post in database
	error := repo.CreatePost(post, nil)
	assert.NoError(t, error)

	// Call the method
	err := repo.DeletePost(post.ID)

	// Assertions
	assert.NoError(t, err)
}
