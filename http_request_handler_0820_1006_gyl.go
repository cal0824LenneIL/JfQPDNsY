// 代码生成时间: 2025-08-20 10:06:12
package main

import (
    "fmt"
    "net/http"
    "log"

    // Import the Echo framework
    "github.com/labstack/echo/v4"
)

// setupRoutes sets up the routes for the HTTP server
func setupRoutes(e *echo.Echo) {
    // Define a route for the root path
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // Define a route for a path that returns a JSON response
    e.GET("/json", func(c echo.Context) error {
        // Return a simple JSON response
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, World!",
        })
    })
}

func main() {
    // Create a new Echo instance
    e := echo.New()

    // Setup routes
    setupRoutes(e)

    // Start the HTTP server
    // The server will listen on port 8080 by default
    if err := e.Start(":8080"); err != nil && err != echo.ErrServerClosed {
        // Handle error
        log.Fatalf("HTTP server failed to start: %v", err)
    }
}
