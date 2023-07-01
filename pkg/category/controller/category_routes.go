package controller

import (
	"github.com/labstack/echo/v4"

	"restful_blog/pkg/middlewares"
)

func CategoryRoutes(e *echo.Group, categoryController CategoryController) {
	e.GET("/category", categoryController.AllCategory, middlewares.Auth)
	e.GET("/category/:id", categoryController.CategoryById, middlewares.Auth)
	e.POST("/category", categoryController.CreateCategory, middlewares.Auth)
	e.PUT("/category/:id", categoryController.UpdateCategory, middlewares.Auth)
	e.DELETE("/category/:id", categoryController.DeleteCategory, middlewares.Auth)
}
