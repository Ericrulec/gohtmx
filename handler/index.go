package handler

import (
	"github.com/Ericrulec/gohtmx/view/pages"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}
type PostsHandler struct{}
type ContentHandler struct{}

func (h *IndexHandler) GetIndex(c echo.Context) error {
	return render(c, pages.HomePage())
}
func (h *PostsHandler) GetPosts(c echo.Context) error {
	return render(c, pages.PostsPage())
}
func (h *ContentHandler) GetPasswordsPage(c echo.Context) error {
	return render(c, pages.PasswordsPage())
}
