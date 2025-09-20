// 代码生成时间: 2025-09-21 07:09:10
 * Features:
 * - Scrapes content from a given URL and returns the HTML.
 * - Implements error handling and logging.
 * - Follows Go best practices for code structure and error handling.
 */

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "github.com/labstack/echo"
)

// ScrapeContent is a handler function that fetches the HTML content from a given URL.
func ScrapeContent(c echo.Context) error {
    url := c.QueryParam("url") // Retrieve the URL from query parameters.
    if url == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "URL parameter is required")
    }

    // Send a GET request to the specified URL.
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error fetching URL: %s
", err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch content")
    }
    defer resp.Body.Close()

    // Read the body of the response.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading response body: %s
", err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read content")
    }

    // Return the HTML content as a response.
    return c.HTMLBlob(http.StatusOK, body)
}

func main() {
    e := echo.New()
    e.GET("/scrape", ScrapeContent) // Register the handler for the /scrape endpoint.

    // Start the Echo server.
    e.Logger.Fatal(e.Start(":8080"))
}
