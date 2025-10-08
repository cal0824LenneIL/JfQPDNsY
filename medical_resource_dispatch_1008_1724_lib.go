// 代码生成时间: 2025-10-08 17:24:46
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// Resource represents a medical resource with its id and name.
type Resource struct {
    ID   int    "json:"id" example:1"
    Name string "json:"name" example:"Ambulance"
}

// ResourceHandler is the handler for the resource dispatching service.
type ResourceHandler struct {
    // Add fields if needed for dependency injection
}

// NewResourceHandler creates a new ResourceHandler.
func NewResourceHandler() *ResourceHandler {
    return &ResourceHandler{}
}

// AddResource adds a new medical resource to the system.
func (h *ResourceHandler) AddResource(c echo.Context) error {
    // Assuming the request payload is a JSON with the resource details.
    var resource Resource
    if err := c.Bind(&resource); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid resource data")
    }

    // Add logic to add the resource to the system
    // For simplicity, we're just logging the resource addition here
    c.Logger().Infof("Adding resource: %+v", resource)

    // Return a success response
    return c.JSON(http.StatusOK, resource)
}

// ListResources lists all the medical resources.
func (h *ResourceHandler) ListResources(c echo.Context) error {
    // Add logic to list resources from the system
    // For simplicity, we're returning a hardcoded list
    var resources = []Resource{
        {ID: 1, Name: "Ambulance"},
        {ID: 2, Name: "Hospital Bed"},
        {ID: 3, Name: "Ventilator"},
    }

    return c.JSON(http.StatusOK, resources)
}

func main() {
    e := echo.New()
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
       -AllowOrigins: []string{"*"},
       AllowMethods: []string{"GET", "POST"},
       AllowHeaders: []string{"Content-Type"},
    }))

    // Register the resource handler with the Echo instance
    resHandler := NewResourceHandler()
    e.POST("/resources", resHandler.AddResource)
    e.GET("/resources", resHandler.ListResources)

    // Start the Echo server
    e.Logger.Fatal(e.Start(":" + "8080"))
}
