package handler

import (
	"github.com/Ericrulec/gohtmx/view/pages"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}
type PostsHandler struct{}
type ContentHandler struct{}

func (h *IndexHandler) GetIndex(c echo.Context) error {
	return pages.HomePage().Render(c.Request().Context(), c.Response())
}
func (h *PostsHandler) GetPosts(c echo.Context) error {
	return pages.PostsPage().Render(c.Request().Context(), c.Response())
}
func (h *ContentHandler) GetPasswordsPage(c echo.Context) error {
	return pages.PasswordsPage().Render(c.Request().Context(), c.Response())
}
