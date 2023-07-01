package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"restful_blog/pkg/post/model"
	"restful_blog/pkg/post/service"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var validate *validator.Validate

type PostController struct {
	PostService service.PostServiceInterface
}

func NewPostController(sp service.PostServiceInterface) *PostController {
	return &PostController{
		PostService: sp,
	}
}

func init() {
	validate = validator.New()
}

func (h *PostController) PostById(c echo.Context) error {
	id := c.Param("id")
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: "Invalid Category ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	post, err := h.PostService.PostByID(postID)
	if err != nil {
		response := Response{
			Message: "Invalid Category ID",
		}
		return c.JSON(http.StatusNotFound, response)
	}
	response := Response{
		Message: "Success",
		Data:    post,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *PostController) AllPost(c echo.Context) error {
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")

	posts, err := h.PostService.AllPosts(pageStr, perPageStr)

	if err != nil {
		response := Response{
			Message: "Invalid Category ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Data:    posts,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *PostController) CreatePost(c echo.Context) error {
	post := new(model.Post)

	if err := c.Bind(post); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := validate.Struct(post); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	// Create the post
	if err := h.PostService.CreatePost(post, post.CategoryIDs); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Data:    post,
		Message: "Success",
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *PostController) UpdatePost(c echo.Context) error {
	id := c.Param("id")
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	// Convert the id from string to int64
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: "Invalid Post ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	// Set the ID of the post to be updated
	post.ID = postID

	if err := h.PostService.UpdatePost(post, post.CategoryIDs); err != nil {
		response := Response{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Data:    post,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *PostController) DeletePost(c echo.Context) error {
	id := c.Param("id")

	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := Response{
			Message: "Invalid Post ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.PostService.DeletePost(postID); err != nil {
		response := Response{
			Message: "Invalid Post ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := Response{
		Message: "Post Deleted Successfully",
	}
	return c.JSON(http.StatusNoContent, response)
}
