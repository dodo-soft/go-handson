package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	message := ""
	e.POST("memo", func(c echo.Context) error {
		message = c.FormValue("message")
		return c.String(200, message)
	})
	e.GET("memo", func(c echo.Context) error {
		return c.String(200, message)
	})
	e.Start(":8080")
}
