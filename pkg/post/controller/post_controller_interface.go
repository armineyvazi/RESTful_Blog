package controller

import (
	"github.com/labstack/echo/v4"
)

type PostControllerInterface interface {
	AllPost(c echo.Context) error
	PostById(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}
