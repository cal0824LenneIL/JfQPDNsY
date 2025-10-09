// 代码生成时间: 2025-10-09 21:27:43
// color_selector.go

package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// Color represents the structure for a color
type Color struct {
    Name    string `json:"name"`
    HexCode string `json:"hex_code"`
}

// ColorHandler handles the request for the color selector
func ColorHandler(c echo.Context) error {
    // Define a slice of colors for demonstration purposes
    colors := []Color{
        {
            Name:    "Red",
            HexCode: "#FF0000",
        },
        {
            Name:    "Green",
            HexCode: "#00FF00",
        },
        {
            Name:    "Blue",
            HexCode: "#0000FF",
        },
    }

    // Return the slice of colors as JSON
    return c.JSON(http.StatusOK, colors)
}

func main() {
    // Create a new Echo instance
    e := echo.New()

    // Define the route for the color selector
    e.GET("/colors", ColorHandler)

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
