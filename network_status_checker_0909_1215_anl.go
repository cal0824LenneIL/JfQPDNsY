// 代码生成时间: 2025-09-09 12:15:47
 * This program periodically checks the network connection status and returns the result.
 *
 * Features:
 * - Clear code structure for easy understanding.
 * - Appropriate error handling.
 * - Necessary comments and documentations.
 * - Adherence to Go best practices.
 * - Ensuring code maintainability and extensibility.
 */

package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/labstack/echo"
)

// Constants for the default ping URL and frequency
const (
    defaultPingURL = "http://www.google.com"
    pingFrequency  = 10 * time.Second // 10 seconds
)

// NetworkChecker is a struct that holds the Echo instance and ping URL
type NetworkChecker struct {
    e    *echo.Echo
    url  string
}

// NewNetworkChecker returns a new instance of NetworkChecker
func NewNetworkChecker(url string) *NetworkChecker {
    return &NetworkChecker{
        e:    echo.New(),
        url: url,
    }
}

// Start starts the Echo server with the network status checking endpoint
func (nc *NetworkChecker) Start() {
    nc.e.GET("/check", nc.checkNetworkStatus)
    log.Fatal(nc.e.Start(":8080"))
}

// checkNetworkStatus checks the network connection status by pinging the URL
func (nc *NetworkChecker) checkNetworkStatus(c echo.Context) error {
    result, err := pingURL(nc.url)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "status":  "error",
            "message": err.Error(),
        })
    }
    return c.JSON(http.StatusOK, echo.Map{
        "status":  "ok",
        "message": result,
    })
}

// pingURL checks if a URL is reachable by sending a HEAD request
func pingURL(url string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
    if err != nil {
        return "", err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusOK {
        return fmt.Sprintf("URL %s is reachable", url), nil
    }
    return "", fmt.Errorf("URL %s is not reachable, status code: %d", url, resp.StatusCode)
}

func main() {
    // Create a new network checker instance with the default ping URL
    nc := NewNetworkChecker(defaultPingURL)
    
    // Start the Echo server
    nc.Start()
}
