package akademy

import (
	"net/http"

	"github.com/labstack/echo"
)

func fetchName(c echo.Context) error {
	return c.JSON(http.StatusOK, "Alwin here")
}
