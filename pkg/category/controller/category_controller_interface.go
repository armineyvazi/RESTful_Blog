package controller

import (
	"github.com/labstack/echo/v4"
)

type CategoryControllerInterface interface {
	GetCategoryById(c echo.Context) error
	GetAllCategory(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}
