package routes

import (
	"go-scaled-api/services/api/handlers"

	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/domains/:domainName", handlers.GetDomainHandler)
	e.PUT("/event/:domainName/delivered", handlers.DeliveredEventHandler)
	e.PUT("/event/:domainName/bounced", handlers.BouncedEventHandler)
	e.GET("/test", handlers.SanityTest)
}
