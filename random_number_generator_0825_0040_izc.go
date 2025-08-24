// 代码生成时间: 2025-08-25 00:40:59
package main

import (
    "crypto/rand"
    "encoding/binary"
    "math/big"
    "net/http"
    "os"
    "github.com/labstack/echo"
)

// RandomNumberGenerator handles the HTTP request for generating random numbers.
func RandomNumberGenerator(c echo.Context) error {
    // Generate a random number between 1 and 100
    randomNumber, err := generateRandomNumber(1, 100)
    if err != nil {
        // Return a 500 Internal Server Error with a message if an error occurs
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Failed to generate random number",
            "message": err.Error(),
        })
    }
    // Return the generated random number in the response
    return c.JSON(http.StatusOK, echo.Map{
        "randomNumber": randomNumber,
    })
}

// generateRandomNumber generates a random number between the given min and max values.
// It returns the generated number and any error that may occur.
func generateRandomNumber(min, max int) (int, error) {
    // Use crypto/rand for a cryptographically secure random number
    randomBytes := make([]byte, 8)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return 0, err
    }
    // Convert the bytes to a big.Int number
    randomNumber, err := big.NewInt(0).SetBytes(randomBytes)
    if err != nil {
        return 0, err
    }
    // Scale down the number to the desired range [min, max]
    scaledNumber := new(big.Int)
    scaledNumber.Mul(randomNumber, big.NewInt(int64(max-min+1)))
    scaledNumber.Add(scaledNumber, big.NewInt(int64(min-1)))
    // Convert the result to an int
    result := scaledNumber.Int64() + 1 // +1 because we want the range [min, max], not [min-1, max-1]
    if result < int64(min) || result > int64(max) {
        return 0, fmt.Errorf("random number out of range: %d", result)
    }
    return int(result), nil
}

func main() {
    e := echo.New()
    e.GET("/random", RandomNumberGenerator)
    
    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}