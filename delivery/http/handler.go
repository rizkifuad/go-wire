package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizkix/wired/controller"
)

type Handler struct {
	Controller controller.Controller
}

func New(c controller.Controller) *echo.Echo {
	e := echo.New()

	handler := Handler{
		Controller: c,
	}

	e.GET("/ping", handler.Ping)
	e.GET("/repo", handler.Get)

	return e
}

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "success"})
}

func (h *Handler) Get(c echo.Context) error {
	a := h.Controller.Get("a")
	return c.JSON(http.StatusOK, a)
}
