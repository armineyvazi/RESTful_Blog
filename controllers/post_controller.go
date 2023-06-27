package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"restful_blog/models"
	"restful_blog/service"
)

var validate *validator.Validate

type PostController struct {
	PostService service.PostServiceInterface
}

func NewPostController() *PostController {
	postService := service.NewPostService()
	return &PostController{
		PostService: postService,
	}
}

func init() {
	validate = validator.New()
}

func (h *PostController) Get(c echo.Context) error {
	id := c.Param("id")

	//with id
	if id != "" {
		postID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"Message": "Invalid Post ID",
			})
		}
		post, err := h.PostService.GetPostByID(postID)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, post)
	}
	// without id
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")

	posts, err := h.PostService.GetAllPosts(pageStr, perPageStr)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (h *PostController) Post(c echo.Context) error {
	post := new(models.Post)

	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create the post
	if err := h.PostService.CreatePost(post, post.CategoryIDs); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}

func (h *PostController) Put(c echo.Context) error {
	id := c.Param("id")
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Convert the id from string to int64
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid Post ID",
		})
	}

	// Set the ID of the post to be updated
	post.ID = postID

	if err := h.PostService.UpdatePost(post, post.CategoryIDs); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostController) Delete(c echo.Context) error {
	id := c.Param("id")

	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid Post ID",
		})
	}

	if err := h.PostService.DeletePost(postID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, map[string]string{
		"Message": "Post Deleted Successfully",
	})
}
