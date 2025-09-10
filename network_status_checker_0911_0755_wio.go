// 代码生成时间: 2025-09-11 07:55:02
 * Features:
 * - Checks the connection status of a given URL.
 * - Handles errors appropriately.
 * - Follows Go best practices for maintainability and scalability.
 */

package main

import (
    "net/http"
    "time"
    "github.com/labstack/echo"
    "log"
)

// NetworkStatusResponse is the response structure for network status checks.
type NetworkStatusResponse struct {
    URL      string `json:"url"`
    Status   string `json:"status"`
    Duration string `json:"duration"`
    Error    string `json:"error,omitempty"`
}

// checkNetworkStatus checks the network connection status of a given URL and returns the result.
func checkNetworkStatus(url string) (NetworkStatusResponse, error) {
    client := &http.Client{
        Timeout: 5 * time.Second,
    }
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return NetworkStatusResponse{}, err
    }
    start := time.Now()
    resp, err := client.Do(req)
    duration := time.Since(start).String()
    if err != nil {
        return NetworkStatusResponse{URL: url, Status: "Failed", Duration: duration, Error: err.Error()}, nil
    }
    defer resp.Body.Close()

    return NetworkStatusResponse{
        URL:      url,
        Status:   "Success",
        Duration: duration,
    }, nil
}

func main() {
    e := echo.New()
    
    // Define the route for checking network status.
    e.GET("/check", func(c echo.Context) error {
        url := c.QueryParam("url")
        if url == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "URL parameter is required")
        }
        result, err := checkNetworkStatus(url)
        if err != nil {
            result.Error = err.Error()
        }
        return c.JSON(http.StatusOK, result)
    })
    
    // Start the server.
    e.Logger.Fatal(e.Start(":8080"))
}
