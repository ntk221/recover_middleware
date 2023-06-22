package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Recover middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", helloHandler)

	// Start server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(c echo.Context) error {
	// Simulate a panic
	panic("Something went wrong!")

	// This line will never be executed
	return c.String(http.StatusOK, "Hello, World!")
}

