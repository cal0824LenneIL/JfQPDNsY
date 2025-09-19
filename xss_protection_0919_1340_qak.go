// 代码生成时间: 2025-09-19 13:40:45
package main

import (
    "net/http"
    "strings"
    "html"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// XSSProtectionMiddleware is a middleware function that sanitizes
// incoming request data to prevent XSS attacks.
func XSSProtectionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Get the raw query parameters
        query := c.Request().URL.Query()

        // Iterate over each query parameter and sanitize it
        for key, values := range query {
            for i, value := range values {
                query[key][i] = sanitizeInput(value)
            }
        }

        // Create a new URL with sanitized query parameters
        c.Request().URL.RawQuery = query.Encode()

        // Continue to the next middleware
        return next(c)
    }
}

// sanitizeInput sanitizes the input to prevent XSS attacks by
// converting any HTML characters to their corresponding HTML entities.
func sanitizeInput(input string) string {
    return html.EscapeString(input)
}

func main() {
    e := echo.New()

    // Register the XSS protection middleware
    e.Use(middleware.Recover())
    e.Use(XSSProtectionMiddleware)

    // Define a route for demonstration purposes
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
