package akademy

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/name", fetchName)
}
