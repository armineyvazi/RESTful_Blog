package service

// Test_Post_Service_Get_Post_By_Id test post service GetPostById
//func Test_Post_Service_Get_Post_By_Id(t *testing.T) {
//	// create mock repository
//	mockRepo := &repositories.MockPostRepository{}
//
//	// set up the expected response from the mock repository
//	expectedPost := &model.Post{
//		ID:          101,
//		Title:       "Test 1",
//		Text:        "Test 1",
//		CategoryIDs: nil,
//	}
//
//	// Mock post repository
//	mockRepo.GetPostsByIDFn = func(postsId int64) (*model.Post, error) {
//		return expectedPost, nil
//	}
//
//	// create a new PostService instance with the mock post repository
//	s := &PostService{
//		PostRepository: mockRepo,
//	}
//
//	//call the method
//	post, err := s.GetPostByID(expectedPost.ID)
//
//	//assertion
//	assert.NoError(t, err)
//	assert.Equal(t, expectedPost, post)
//}
//
//// Test_Post_Service_Get_All_Post  test post service GetAllPost method
//func Test_Post_Service_Get_All_Post(t *testing.T) {
//	// Create a mock repository
//	mockRepo := &repositories.MockPostRepository{}
//
//	// Set up the expected response from the mock post repository
//	expectedPost := []*model.Post{
//		{
//			ID:          101,
//			Title:       "Test 1",
//			Text:        "Test 1",
//			CategoryIDs: nil,
//		},
//		{
//			ID:          102,
//			Title:       "Test 2",
//			Text:        "Test 2",
//			CategoryIDs: nil,
//		},
//	}
//
//	mockRepo.GetAllPostsFn = func(page, perPage int) ([]*model.Post, error) {
//		return expectedPost, nil
//	}
//
//	// Create a new PostService instance with the mock post repository
//	s := &PostService{
//		PostRepository: mockRepo,
//	}
//
//	// Call the method
//	post, err := s.GetAllPosts("1", "10")
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, expectedPost, post)
//}
//
//// Test_Post_Service_Create_Post test post service CreatePost method
//func Test_Post_Service_Create_Post(t *testing.T) {
//	// Create a mock repository
//	mockRepo := &repositories.MockPostRepository{}
//
//	// Set up the expected response from the mock repository
//	mockRepo.CreatePostsFn = func(post *model.Post, categoryIDs []int64) error {
//		post.ID = 1 // Assign the expected ID
//		return nil
//	}
//
//	// Create a new PostService  instance with the mock repository
//	s := &PostService{
//		PostRepository: mockRepo,
//	}
//
//	// Create a new post
//	post := &model.Post{
//		ID:          101,
//		Title:       "Test 1",
//		Text:        "Test 1",
//		CategoryIDs: nil,
//	}
//	// Call the method
//	err := s.CreatePost(post, nil)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, int64(1), post.ID)
//}
//
//// Test_Post_Service_Update_Post test post service UpdateService method
//func Test_Post_Service_Update_Post(t *testing.T) {
//	// Create a mock repository
//	mockRepo := &repositories.MockPostRepository{}
//
//	// Set up the expected response from the mock  post repository
//	mockRepo.UpdatePostFn = func(post *model.Post, category []int64) error {
//		return nil
//	}
//
//	// Create a new PostService instance with the mock post repository
//	s := &PostService{
//		PostRepository: mockRepo,
//	}
//
//	// Create a post to update
//	post := &model.Post{
//		ID:          101,
//		Title:       "Test 1",
//		Text:        "Test 1",
//		CategoryIDs: nil,
//	}
//
//	// Call the method
//	err := s.UpdatePost(post, nil)
//
//	// Assertions
//	assert.NoError(t, err)
//}
//
//// Test_Post_Service_Delete_Post test post service DeletePost method
//func Test_Post_Service_Delete_Post(t *testing.T) {
//	// Create a mock repository
//	mockRepo := &repositories.MockPostRepository{}
//
//	// Set up the expected response from the mock post repository
//	mockRepo.DeletePostFn = func(categoryID int64) error {
//		return nil
//	}
//
//	// Create a new CategoryService instance with the mock post repository
//	s := &PostService{
//		PostRepository: mockRepo,
//	}
//
//	// Call the method
//	err := s.DeletePost(1)
//
//	// Assertions
//	assert.NoError(t, err)
//}
