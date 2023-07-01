package repository

//// Test_Category_Repository_Get_By_Id test category repository  GetById  method
//func Test_Category_Repository_Get_By_Id(t *testing.T) {
//	// Create a new CategoryRepositoryImpl instance
//	repo := NewCategoryRepository()
//
//	// Create a category to be retrieved
//	expectedCategory := &model.Category{
//		ID:   102,
//		Name: "Category 1",
//	}
//	error := repo.CreateCategory(expectedCategory)
//	assert.NoError(t, error)
//
//	// Call the method
//	category, err := repo.GetCategoryByID(expectedCategory.ID)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, expectedCategory, category)
//
//	// clean database
//	err = repo.DeleteCategory(expectedCategory.ID)
//	assert.NoError(t, err)
//	fmt.Println("clean database....")
//
//}
//
//// Test_Category_Repository_Get_All test category repository GetAll method
//func Test_Category_Repository_Get_All(t *testing.T) {
//
//	// Create a new CategoryRepositoryImpl instance
//	repo := NewCategoryRepository()
//
//	// Create some categories in the database
//	expectedCategories1 := &model.Category{
//		ID:   105,
//		Name: "Category 1",
//	}
//
//	expectedCategories2 := &model.Category{
//		ID:   106,
//		Name: "Category 1",
//	}
//
//	// create expected 1 category
//	error := repo.CreateCategory(expectedCategories1)
//	assert.NoError(t, error)
//
//	// create expected 2 category
//	error = repo.CreateCategory(expectedCategories2)
//	assert.NoError(t, error)
//
//	// Call the method
//	categories, err := repo.GetAllCategories(1, 10)
//
//	fmt.Println("cate", categories[0].ID)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, expectedCategories1.ID, categories[0].ID)
//	assert.NoError(t, err)
//	assert.Equal(t, expectedCategories2.ID, categories[1].ID)
//
//	// clean database
//	err = repo.DeleteCategory(expectedCategories1.ID)
//	assert.NoError(t, err)
//	err = repo.DeleteCategory(expectedCategories2.ID)
//	assert.NoError(t, err)
//}
//
//// Test_Category_Repository_Create_Category test category repository CreateCategory method
//func Test_Category_Repository_Create_Category(t *testing.T) {
//	// Create a new CategoryRepositoryImpl instance
//	repo := NewCategoryRepository()
//	// Create a new category
//
//	category := &model.Category{
//		Name: "New Category",
//	}
//
//	// Call the method
//	err := repo.CreateCategory(category)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.NotZero(t, category.ID)
//
//	err = repo.DeleteCategory(category.ID)
//	assert.NoError(t, err)
//}
//
//// Test_Update_Category test category repository UpdateCategory method
//func Test_Category_Repository_Update_Category(t *testing.T) {
//	// Create a new CategoryRepositoryImpl instance
//	repo := NewCategoryRepository()
//
//	// Create a category to update
//	category := &model.Category{
//		Name: "Updated Category",
//	}
//
//	error := repo.CreateCategory(category)
//	assert.NoError(t, error)
//
//	// Call the method
//	err := repo.UpdateCategory(category)
//
//	// Assertions
//	assert.NoError(t, err)
//
//	err = repo.DeleteCategory(category.ID)
//	assert.NoError(t, err)
//
//}
//
//// Test_Category_Repository_DeleteCategory test category repository DeleteCategory method
//func Test_Category_Repository_DeleteCategory(t *testing.T) {
//	// Create a new CategoryRepositoryImpl instance
//	repo := NewCategoryRepository()
//
//	// Create a category to delete
//	category := &model.Category{
//		ID:   101,
//		Name: "Category to delete",
//	}
//
//	error := repo.CreateCategory(category)
//	assert.NoError(t, error)
//
//	// Call the method
//	err := repo.DeleteCategory(category.ID)
//
//	// Assertions
//	assert.NoError(t, err)
//}
