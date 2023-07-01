package controller

//import (
//	"encoding/json"
//	"net/http"
//	"net/http/httptest"
//	_ "strconv"
//	"strings"
//	"testing"
//
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//
//	service_mock "restful_blog/apps/blog/mock/service"
//
//	"restful_blog/pkg/category/model"
//	"restful_blog/pkg/routes"
//)
//
//var e *echo.Echo
//
//func init() {
//	e = echo.New()
//	routes.Router(e)
//}
//
//// Test_category_controller_Get_By_Id test category controller method by id
//func Test_category_controller_Get_By_Id(t *testing.T) {
//	// Mock category service
//	mockCategoryService := &service_mock.MockCategoryService{}
//
//	//pass mock category service to category controller
//	controller := &CategoryController{
//		CategoryService: mockCategoryService,
//	}
//
//	// Prepare the request with a valid ID
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/1?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	// Call the Get method in the category controller
//	err := controller.Get(c)
//
//	// Check for any errors
//	if err != nil {
//		t.Errorf("Error retrieving category: %s", err.Error())
//	}
//	// Check the response status code
//	if rec.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
//	}
//}
//
//// Test_category_controller_Get_By_Id_Not_Valid_ID test category controller get not valid id
//func Test_category_controller_Get_By_Id_Not_Valid_ID(t *testing.T) {
//	// Mock CategoryService
//	mockCategoryService := &service_mock.MockCategoryService{}
//	controller := &CategoryController{
//		CategoryService: mockCategoryService,
//	}
//
//	// Prepare the request with a not valid ID
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/asdasd?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	e.ServeHTTP(rec, req)
//
//	// Call the Get method on the category controller
//	err := controller.Get(c)
//
//	// Check for any errors
//	if err != nil {
//		t.Errorf("Error retrieving category: %s", err.Error())
//	}
//	// Check the response status code
//	if rec.Code != http.StatusBadRequest {
//		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
//	}
//}
//
//// Test_category_controller_Get_All test category controller  GetAll method
//func Test_category_controller_Get_All(t *testing.T) {
//	// Mock CategoryService
//	mockCategoryService := &service_mock.MockCategoryService{}
//
//	//mockCategoryRepository := repository.MockCategoryRepository{}
//	controller := &CategoryController{
//		CategoryService: mockCategoryService,
//	}
//
//	// Prepare the request with a valid ID
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/category?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	// Call the Get method on the category controller
//	err := controller.Get(c)
//
//	// Check for any errors
//	if err != nil {
//		t.Errorf("Error retrieving category: %s", err.Error())
//	}
//	// Check the response status code
//	if rec.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
//	}
//}
//
//// Test_put_category test category controller put method
//func Test_put_category(t *testing.T) {
//	// Create a mock service
//	mockService := &service_mock.MockCategoryService{}
//
//	// Set up the expected response from the mock service
//	expectedCategory := &model.Category{
//		ID:   1,
//		Name: "Updated Category",
//	}
//	mockService.UpdateCategoryFunc = func(category *model.Category) error {
//		return nil
//	}
//
//	// Create a new CategoryController instance with the mock service
//	controller := &CategoryController{
//		CategoryService: mockService,
//	}
//
//	// Create a request payload
//	requestBody := `{"name": "Updated Category"}`
//	req := httptest.NewRequest(http.MethodPut, "/api/v1/category/1?api_key=armin", strings.NewReader(requestBody))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	// Call the handler method
//	err := controller.Put(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//	var responseCategory model.Category
//	err = json.Unmarshal(rec.Body.Bytes(), &responseCategory)
//	assert.NoError(t, err)
//	assert.Equal(t, expectedCategory, &responseCategory)
//}
//
//// Test_delete_category test delete category controller method
//func Test_delete_category(t *testing.T) {
//	// Create a mock service
//	mockService := &service_mock.MockCategoryService{}
//
//	// Set up the expected response from the mock service
//	mockService.DeleteCategoryFunc = func(categoryID int64) error {
//		return nil
//	}
//
//	// Create a new CategoryController instance with the mock service
//	controller := &CategoryController{
//		CategoryService: mockService,
//	}
//
//	// Create a request
//	req := httptest.NewRequest(http.MethodDelete, "/api/v1/category/1?api_key=armin", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	// Call the handler method
//	err := controller.Delete(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusNoContent, rec.Code)
//}
//
//func TestPostCategory(t *testing.T) {
//	// Create a new Echo instance
//	e := echo.New()
//
//	// Create a mock service
//	mockService := &service_mock.MockCategoryService{}
//
//	// Set up the expected response from the mock service
//	expectedCategory := &model.Category{
//		ID:   1,
//		Name: "New Category",
//	}
//	mockService.CreateCategoryFunc = func(category *model.Category) error {
//		category.ID = expectedCategory.ID // Assign the expected ID
//		return nil
//	}
//
//	// Create a new CategoryController instance with the mock service
//	controller := &CategoryController{
//		CategoryService: mockService,
//	}
//
//	// Create a request payload
//	requestBody := `{"name": "New Category"}`
//	req := httptest.NewRequest(http.MethodPost, "/api/v1/category?api_key=armin", strings.NewReader(requestBody))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//
//	// Call the handler method
//	err := controller.Post(c)
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusCreated, rec.Code)
//
//	var responseCategory model.Category
//	err = json.Unmarshal(rec.Body.Bytes(), &responseCategory)
//	assert.NoError(t, err)
//	assert.Equal(t, expectedCategory, &responseCategory)
//}
