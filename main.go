package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/status", func(c echo.Context) error {
		status := "InProgress"
		now := time.Now()
		lastDigit := now.Second() % 10
		if lastDigit == 0 {
			status = "Success"
		}

		return c.String(http.StatusOK, status)
	})
	e.Logger.Panic(e.Start(":1323"))
}
