package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "restful_blog/apps/blog/docs"
	"restful_blog/controllers"
	"restful_blog/middlewares"
)

func Router(e *echo.Echo) {
	e.HideBanner = true
	e.Debug = true

	//swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Group routes with API version prefix
	v1 := e.Group("/api/v1")
	{

		// Create instances of the controllers
		postController := controllers.NewPostController() // Create a new instance of PostController
		categoryController := controllers.NewCategoryController()

		//post api
		{
			v1.GET("/posts", postController.Get, middlewares.Auth)
			v1.GET("/posts/:id", postController.Get, middlewares.Auth)
			v1.POST("/posts", postController.Post, middlewares.Auth)
			v1.PUT("/posts/:id", postController.Put, middlewares.Auth)
			v1.DELETE("/posts/:id", postController.Delete, middlewares.Auth)
		}

		//category api
		{
			v1.GET("/category", categoryController.Get, middlewares.Auth)
			v1.GET("/category/:id", categoryController.Get, middlewares.Auth)
			v1.POST("/category", categoryController.Post, middlewares.Auth)
			v1.PUT("/category/:id", categoryController.Put, middlewares.Auth)
			v1.DELETE("/category/:id", categoryController.Delete, middlewares.Auth)
		}
	}
}
