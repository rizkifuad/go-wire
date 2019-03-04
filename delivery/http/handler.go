package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizkix/wired/controller"
)

type Handler struct {
	Controller controller.Controller
	Instance   *echo.Echo
}

func New(c controller.Controller) Handler {
	e := echo.New()

	handler := Handler{
		Controller: c,
	}

	e.GET("/ping", handler.Ping)
	e.GET("/repo", handler.Get)

	return Handler{Instance: e}
}

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "success"})
}

func (h *Handler) Get(c echo.Context) error {
	a := h.Controller.Get("a")
	return c.JSON(http.StatusOK, a)
}
