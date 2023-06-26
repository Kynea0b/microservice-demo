package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("API Gateway Start!!")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/draw", Draw)
	e.GET("/histories/:user_id", GetHistories)
	e.GET("/inventories/:user_id", GetInventories)
	e.GET("/test", func(c echo.Context) error {
		return c.String(200, "test")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
