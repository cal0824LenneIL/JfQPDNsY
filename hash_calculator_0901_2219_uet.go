// 代码生成时间: 2025-09-01 22:19:55
 * This tool allows users to calculate the hash of a given string input.
 *
 * @author Your Name
 * @date 2023-04-20
 */

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// StartHashCalculator starts the Echo web server for the hash calculator tool.
func StartHashCalculator() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Route
    e.POST("/hash", hashHandler)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}

// hashHandler handles the POST request to calculate the hash.
func hashHandler(c echo.Context) error {
    input := c.QueryParam("input")
    if input == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Input parameter 'input' is required.")
    }

    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // Response
    return c.JSON(http.StatusOK, map[string]string{
        "input": input,
        "hash": hashString,
    })
}

func main() {
    StartHashCalculator()
}
