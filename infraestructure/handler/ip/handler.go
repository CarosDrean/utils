package ip

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) ipClient(c echo.Context) error {
	return c.JSON(http.StatusOK, c.RealIP())
}
