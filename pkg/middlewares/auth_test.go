package middlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_unauthorized(t *testing.T) {
	var httpErr *echo.HTTPError
	var ok bool

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/category?api_key=not_valid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	}

	mw := Auth(handler)
	err := mw(c)

	// Get the error code from the error, if applicable
	if ok = errors.As(err, &httpErr); ok {
		assert.Equal(t, http.StatusUnauthorized, httpErr.Code)
	} else {
		t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, rec.Code)
	}

}
