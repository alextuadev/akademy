package server

import (
	"github.com/labstack/echo"
)

func createRoutes(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	createAdminRoutes(adminGroup)

	userGroup := e.Group("/users")
	createUserRoutes(userGroup)

	deptGroup := e.Group("/dept")
	createDeptRoutes(deptGroup)
}

func serveUI(e *echo.Echo) {
	e.Static("/", "ui/webapp")
}

func createAdminRoutes(g *echo.Group) {
}

func createUserRoutes(g *echo.Group) {

}

func createDeptRoutes(g *echo.Group) {

}
