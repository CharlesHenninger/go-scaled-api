package main

import (
	"go-scaled-api/services/api/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.SetupRoutes(e)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
