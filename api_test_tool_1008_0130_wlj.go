// 代码生成时间: 2025-10-08 01:30:22
 * This tool allows users to send HTTP requests to test different endpoints.
 *
 * @author Your Name
 * @version 1.0.0
 * @since 2023-04-01
 */

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/labstack/echo/v4"
)

// Main function, the entry point of the application
func main() {
    // Create a new Echo instance
    e := echo.New()

    // Define routes
    e.GET("/api", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, World! This is the API testing tool.",
        })
    })

    // Start the Echo server
    e.Server.Addr = ":8080" // Set the server address
    go func() {
        if err := e.Start(e.Server.Addr); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Echo server start failed: %s", err)
        }
    }()

    // Wait for an interrupt signal to gracefully shutdown the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down the server...
")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    if err := e.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
    cancel()
    log.Println("Server exiting...
")
}
