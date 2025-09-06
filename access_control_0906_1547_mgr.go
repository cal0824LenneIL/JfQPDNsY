// 代码生成时间: 2025-09-06 15:47:05
package main

import (
    "context"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

// Role defines different roles with different permissions
type Role string

const (
    ROLE_USER Role = "user"
    ROLE_ADMIN Role = "admin"
)

// RolePermissions defines the permissions for different roles
var RolePermissions = map[Role]map[string]bool{
    ROLE_USER: {
        "GET:/api/public":  true,
    },
    ROLE_ADMIN: {
        "GET:/api/admin":     true,
        "GET:/api/public":    true,
    },
}

// Middleware to check if a user has permission to access the endpoint
func PermissionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        role := c.Get("role").(Role)
        auth := c.Get("auth").(bool)
        
        // If user is not authenticated, return unauthorized
        if !auth {
            return echo.NewHTTPError(http.StatusUnauthorized)
        }
        
        path := c.Path()
        method := c.Request().Method
        
        // Check if the user has permission to access the endpoint
        if !hasPermission(role, path, method) {
            return echo.NewHTTPError(http.StatusForbidden)
        }
        
        return next(c)
    }
}

// Function to check if a user has permission to access a specific endpoint
func hasPermission(role Role, path, method string) bool {
    permissions, exists := RolePermissions[role]
    if !exists {
        return false
    }
    endPoint := method + ":" + path
    return permissions[endPoint]
}

func main() {
    e := echo.New()
    
    // Middleware to enable authentication
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())
    e.Use(middleware.Recover())
    e.Use(middleware.Gzip())

    // Middleware to check permissions
    e.Use(PermissionMiddleware)
    
    // Public API endpoint
    e.GET("/api/public", func(c echo.Context) error {
        return c.String(http.StatusOK, "Public API: Hello, World!")
    })

    // Admin API endpoint
    e.GET("/api/admin", func(c echo.Context) error {
        role := c.Get("role").(Role)
        c.Set("role", ROLE_ADMIN)
        c.Set("auth", true)
        return c.String(http.StatusOK, "Admin API: Hello, Admin!")
    })

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
