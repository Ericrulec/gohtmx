package handler

import (
	"github.com/Ericrulec/gohtmx/view/pages"
	"github.com/labstack/echo"
)

type IndexHandler struct{}

func (h *IndexHandler) GetIndex(c echo.Context) error {
	return pages.HomePage().Render(c.Request().Context(), c.Response())
}
