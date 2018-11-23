package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	message := ""
	// メモ取得
	e.GET("memo", func(c echo.Context) error {
		return c.String(200, message)
	})
	// メモの保存
	// ...
	e.Start(":8080")
}
