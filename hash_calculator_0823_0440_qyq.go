// 代码生成时间: 2025-08-23 04:40:42
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"github.com/labstack/echo/v4"
	"log"
)

// hashCalculatorHandler handles the request for hashing data.
func hashCalculatorHandler(c echo.Context) error {
	input := c.QueryParam("input")
	if input == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Input parameter is required")
	}

	// Calculate the SHA-256 hash of the input string.
	hash := sha256.Sum256([]byte(input))

	// Return the hexadecimal representation of the hash.
	return c.JSON(http.StatusOK, map[string]string{
		"input": input,
		"hash": hex.EncodeToString(hash[:]),
	})
}

func main() {
	e := echo.New()

	// Define a route for the hash calculator.
	e.GET("/hash", hashCalculatorHandler)

	// Start the Echo server.
	e.Logger.Fatal(e.Start(":8080"))
}
