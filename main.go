package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})
	e.Logger.Panic(e.Start(":1323"))
}
