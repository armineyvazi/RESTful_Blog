package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	service_mock "restful_blog/apps/blog/mock/service"
	"restful_blog/controllers"
	"restful_blog/router"
)

var e *echo.Echo

func init() {
	e = echo.New()
	router.Router(e)
}

func Test_unauthorized(t *testing.T) {
	//MockCategoryService
	mockCategoryService := &service_mock.MockCategoryService{}

	controller := &controllers.CategoryController{
		CategoryService: mockCategoryService,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category?api_key=NotValid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	// Call the Get method on the category controller
	err := controller.Get(c)

	// Check for any errors
	if err != nil {
		t.Errorf("Error retrieving category: %s", err.Error())
	}

	// Check the response status code
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, rec.Code)
	}
}
