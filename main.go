package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Status struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4000"},
		AllowMethods: []string{http.MethodGet, http.MethodHead},
	}))

	e.GET("/status", func(c echo.Context) error {
		status := "InProgress"
		now := time.Now()
		lastDigit := now.Second() % 10
		if lastDigit == 0 {
			status = "Success"
		}

		return c.JSON(http.StatusOK, &Status{status})
	})
	e.Logger.Panic(e.Start(":1323"))
}
