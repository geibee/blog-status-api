package main

import (
	"log"
	"net/http"
	"time"

	"blog-status-api/sse"

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

	e.GET("/sse", func(c echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		count := 0
		quit := make(chan bool)

		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case <-ticker.C:
				event := sse.Event{
					Data: []byte("InProgress"),
				}
				if err := event.MarshalTo(w); err != nil {
					return err
				}
				w.Flush()
				if count >= 3 {
					close(quit)
				}
				count++
			case <-quit:
				event := sse.Event{
					Data: []byte("Success"),
				}
				if err := event.MarshalTo(w); err != nil {
					return err
				}
				w.Flush()
				return nil
			}
		}
	})

	e.Logger.Panic(e.Start(":1323"))
}
