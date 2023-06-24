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
	e.GET("/history/:user_id", GetHistories)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
