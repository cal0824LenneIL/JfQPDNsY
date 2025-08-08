// 代码生成时间: 2025-08-08 09:00:00
 * It provides a simple HTTP endpoint to calculate the hash of the input data.
 *
 * @author Your Name
 * @version 1.0
 */

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "net/http"
    "github.com/labstack/echo/v4"
    "log"
)

// HashResponse defines the structure for the hash calculation response.
type HashResponse struct {
    Hash string `json:"hash"`
}

// hashHandler calculates the SHA-256 hash of the provided input and returns it as a response.
func hashHandler(c echo.Context) error {
    input := c.QueryParam("input")
    if input == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Input parameter is required")
    }

    // Calculate the SHA-256 hash of the input.
    hash := sha256.Sum256([]byte(input))
    hashStr := hex.EncodeToString(hash[:])

    // Prepare and return the hash response.
    return c.JSON(http.StatusOK, HashResponse{Hash: hashStr})
}

func main() {
    e := echo.New()
    
    // Define the route for the hash calculation endpoint.
    e.GET("/hash", hashHandler)

    // Start the Echo server.
    log.Printf("Server is running at :8080")
    err := e.Start(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
