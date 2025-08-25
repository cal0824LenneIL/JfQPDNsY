// 代码生成时间: 2025-08-25 22:12:22
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "encoding/json"
    "log"
)

// ApiResponse defines the structure for API responses.
type ApiResponse struct {
    Status  string `json:"status"`
    Data    interface{} `json:"data"`
    Message string `json:"message"`
    Error   string `json:"error"`
}

// StartServer starts the Echo server and listens for incoming requests.
func StartServer() {
    e := echo.New()
    e.GET("/format", formatResponse)
    e.Logger.Fatal(e.Start(":1323"))
}

// formatResponse is the handler function for the /format endpoint.
// It takes the data and message as parameters, and returns a formatted API response.
func formatResponse(c echo.Context) error {
    // Extract parameters from the query string
    data := c.QueryParam("data")
    message := c.QueryParam("message")
    
    // Create an API response
    apiResponse := ApiResponse{
        Status:  "success",
        Data:    data,
        Message: message,
        Error:   "",
    }

    // Serialize the response to JSON
    response, err := json.MarshalIndent(apiResponse, "", "    ")
    if err != nil {
        // Handle JSON marshaling error
        apiResponse.Status = "error"
        apiResponse.Message = "Failed to format response"
        apiResponse.Error = err.Error()
        return c.JSON(http.StatusInternalServerError, apiResponse)
    }

    // Return the formatted JSON response
    return c.Blob(http.StatusOK, "application/json", response)
}

func main() {
    // Start the server
    StartServer()
}
