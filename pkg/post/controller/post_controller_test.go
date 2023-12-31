package controller

//
//import (
//	"encoding/json"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//
//	"restful_blog/apps/blog/mock/service"
//	controllers2 "restful_blog/pkg/category/controller"
//	"restful_blog/pkg/post/model"
//	"restful_blog/pkg/routes"
//)
//
//func init() {
//	controllers2.e = echo.New()
//	routes.Router(controllers2.e)
//
//}
//
//// Test_post_controller_Get_All test post controller get by id method
//func Test_post_controller_Get_All(t *testing.T) {
//	// Create a mock PostService
//	mockPostService := &service.MockPostService{}
//
//	// Create a new PostController with the mock service
//	postController := &PostController{
//		PostService: mockPostService,
//	}
//
//	// Create a request
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := controllers2.e.NewContext(req, rec)
//
//	// Mock the GetPosts function in the PostService
//	mockPostService.GetAllPostFunc = func(pageStr, perPageStr string) ([]*model.Post, error) {
//		// Return a mock list of posts
//		posts := []*model.Post{
//			{ID: 1, Title: "Post 1"},
//			{ID: 2, Title: "Post 2"},
//		}
//		return posts, nil
//	}
//
//	// Call the Get function on the PostController
//	err := postController.Get(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//	expectedPosts := []*model.Post{
//		{ID: 1, Title: "Post 1"},
//		{ID: 2, Title: "Post 2"},
//	}
//	expectedResponse, err := json.Marshal(expectedPosts)
//	assert.NoError(t, err)
//
//	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
//}
//
//// Test_post_controller_Get_By_Id test post controller get by id method
//func Test_post_controller_Get_By_Id(t *testing.T) {
//	// Create a mock PostService
//	mockPostService := &service.MockPostService{}
//
//	// Create a new PostController with the mock service
//	postController := &PostController{
//		PostService: mockPostService,
//	}
//
//	// Create a request
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts/1?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := controllers2.e.NewContext(req, rec)
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	// Mock the GetPostByID function in the PostService
//	mockPostService.GetPostByIDFunc = func(id int64) (*model.Post, error) {
//		// Return a mock post
//		post := &model.Post{
//			ID:    1,
//			Title: "Post 1",
//		}
//		return post, nil
//	}
//
//	// Call the GetByID function on the PostController
//	err := postController.Get(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//	expectedPost := &model.Post{
//		ID:    1,
//		Title: "Post 1",
//	}
//	expectedResponse, err := json.Marshal(expectedPost)
//	assert.NoError(t, err)
//
//	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
//}
//
//// Test_post_controller_Delete test post controller delete method
//func Test_post_controller_Delete(t *testing.T) {
//	// Create a mock PostService
//	mockPostService := &service.MockPostService{}
//
//	// Create a new PostController with the mock service
//	postController := &PostController{
//		PostService: mockPostService,
//	}
//
//	// Create a request
//	req := httptest.NewRequest(http.MethodDelete, "/posts/1", nil)
//	rec := httptest.NewRecorder()
//	c := controllers2.e.NewContext(req, rec)
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	// Mock the DeletePost function in the PostService
//	mockPostService.DeletePostFunc = func(id int64) error {
//		// Perform any necessary assertions on the post ID
//		assert.Equal(t, int64(1), id)
//		return nil
//	}
//
//	// Call the Delete function on the PostController
//	err := postController.Delete(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusNoContent, rec.Code)
//	// it's not pass because there isn't any message
//	//assert.Equal(t, "Post deleted successfully", rec.Body.String())
//}
