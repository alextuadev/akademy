package akademy

import (
	"os"

	"github.com/labstack/echo"
)

// Run runs the akademy server
func Run() {
	var port string
	var portSet bool
	if port, portSet = os.LookupEnv("PORT"); !portSet {
		port = "8080"
	}
	e := echo.New()
	e.Static("/", "dist/ui")
	registerRoutes(e)
	e.Logger.Fatal(e.Start(":" + port))
}
