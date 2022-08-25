package model

import (
	"github.com/labstack/echo/v4"
)

type RouterSpecification struct {
	Config Configuration
	Api    *echo.Echo
}
