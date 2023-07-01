package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "restful_blog/docs"

	"restful_blog/di"
	cc "restful_blog/pkg/category/controller"
	pc "restful_blog/pkg/post/controller"
)

func Router(e *echo.Echo, container *di.Container) {
	//swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Group routes with API version prefix
	v1 := e.Group("/api/v1")

	pc.PostRoutes(v1, *container.PostController)
	cc.CategoryRoutes(v1, *container.CategoryController)

}
