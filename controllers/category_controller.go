package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"restful_blog/models"
	"restful_blog/service"
)

type CategoryController struct {
	CategoryService service.CategoryServiceInterface
}

func NewCategoryController() *CategoryController {
	categoryService := service.NewCategoryService()
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (h *CategoryController) Get(c echo.Context) error {
	id := c.Param("id")

	if id != "" {
		categoryID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"Message": "Invalid Category ID",
			})
		}
		category, err := h.CategoryService.GetCategoryByID(categoryID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"Message": "Record Not Found",
			})
		}
		return c.JSON(http.StatusOK, category)
	}

	// without id
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")
	categories, err := h.CategoryService.GetAllCategory(pageStr, perPageStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Record Not Found",
		})
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryController) Post(c echo.Context) error {
	category := new(models.Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Not Valid Request",
		})
	}

	if err := h.CategoryService.CreateCategory(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Not Valid Request",
		})
	}
	return c.JSON(http.StatusCreated, category)
}

func (h *CategoryController) Put(c echo.Context) error {
	id := c.Param("id")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid Category ID",
		})
	}

	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Not Valid Request",
		})
	}

	category.ID = categoryID

	if err := h.CategoryService.UpdateCategory(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Not Valid Request",
		})
	}
	return c.JSON(http.StatusOK, category)
}

func (h *CategoryController) Delete(c echo.Context) error {
	id := c.Param("id")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid post ID",
		})
	}

	if err := h.CategoryService.DeleteCategory(categoryID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid post ID",
		})
	}
	return c.JSON(http.StatusNoContent, map[string]string{
		"Message": "Category Deleted Successfully",
	})
}
