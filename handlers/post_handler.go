package handlers

//sare handlers bass ERROR return karenge

import (
	"net/http"

	"github.com/C0deNe0/blog-go/domain"
	"github.com/C0deNe0/blog-go/services"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(e *echo.Echo, s services.PostService) {
	h := &PostHandler{service: s}

	posts := e.Group("/posts")
	posts.POST("", h.Create)
	posts.GET("/:id", h.GetById)
	posts.DELETE("/:id", h.Delete)
	posts.PUT("/:id", h.Update)
	posts.GET("", h.List)

}

// create handlers
func (h *PostHandler) Create(c echo.Context) error {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}

	err := h.service.CreatePost(c.Request().Context(), post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Post created successfully"})
}

func (h *PostHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	post, err := h.service.GetPostById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post not found")

	}
	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) List(c echo.Context) error {
	posts, err := h.service.ListPosts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}

	post.Id = id
	err := h.service.UpdatePost(c.Request().Context(), post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (h *PostHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeletePost(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}
