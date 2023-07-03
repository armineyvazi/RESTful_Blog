package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"restful_blog/pkg/category/model"
	"restful_blog/pkg/category/service"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CategoryController struct {
	CategoryService service.CategoryServiceInterface
}

func NewCategoryController(sc service.CategoryServiceInterface) *CategoryController {
	return &CategoryController{
		CategoryService: sc,
	}
}

func (h *CategoryController) CategoryById(c echo.Context) error {
	id := c.Param("id")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: "Invalid Category ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	category, err := h.CategoryService.CategoryByID(categoryID)
	if err != nil {
		response := Response{
			Message: "Record Not Found",
		}
		return c.JSON(http.StatusNotFound, response)
	}
	response := Response{
		Message: "Success",
		Data:    category,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *CategoryController) AllCategory(c echo.Context) error {
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")
	categories, err := h.CategoryService.AllCategory(pageStr, perPageStr)
	if err != nil {
		response := Response{
			Message: "Record Not Found",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := Response{
		Data:    categories,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *CategoryController) CreateCategory(c echo.Context) error {
	category := new(model.Category)

	if err := c.Bind(category); err != nil {
		response := Response{
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	if _, err := h.CategoryService.CreateCategory(category); err != nil {
		response := Response{
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Message: "Success",
		Data:    category,
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *CategoryController) UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: "Invalid Category ID",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	category := new(model.Category)
	if err := c.Bind(category); err != nil {
		response := Response{
			Message: "Not Valid Request",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	category.ID = categoryID

	if _, err := h.CategoryService.UpdateCategory(category); err != nil {
		response := Response{
			Message: "Not Valid Request",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Data:    category,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.CategoryService.DeleteCategory(categoryID); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Message: "Category Deleted Successfully",
	}
	return c.JSON(http.StatusNoContent, response)
}
