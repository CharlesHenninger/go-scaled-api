package handlers

import (
	"github.com/labstack/echo"
)

func SanityTest(c echo.Context) error {
	return c.String(200, "HELLO WORLD")
}
