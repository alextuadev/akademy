package server

import (
	"os"

	"github.com/labstack/echo"
)

// Run runs the akademy server
func Run() {
	e := createEchoServer()
	var akademyServerPort string
	if akademyServerPort = os.Getenv("AKADEMY_SERVER_PORT"); akademyServerPort == "" {
		akademyServerPort = "8080"
	}
	e.Logger.Fatal(e.Start(":" + akademyServerPort))
}

func createEchoServer() *echo.Echo {
	e := echo.New()
	return e
}
