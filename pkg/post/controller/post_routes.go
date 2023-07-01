package controller

import (
	"github.com/labstack/echo/v4"

	"restful_blog/pkg/middlewares"
)

func PostRoutes(e *echo.Group, postController PostController) {
	e.GET("/posts", postController.AllPost, middlewares.Auth)
	e.GET("/posts/:id", postController.PostById, middlewares.Auth)
	e.POST("/posts", postController.CreatePost, middlewares.Auth)
	e.PUT("/posts/:id", postController.UpdatePost, middlewares.Auth)
	e.DELETE("/posts/:id", postController.DeletePost, middlewares.Auth)
}
