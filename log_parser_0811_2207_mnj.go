// 代码生成时间: 2025-08-11 22:07:14
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/labstack/echo"
)

// LogEntry represents a single log entry with timestamp and message
type LogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Message   string    `json:"message"`
}

// parseLogFile reads the contents of a log file and parses it into LogEntry structs
func parseLogFile(filePath string) ([]LogEntry, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var entries []LogEntry
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Assuming log format: "timestamp message"
        parts := strings.Fields(line)
        if len(parts) < 2 {
            continue // Skip lines that don't contain enough parts
        }

        timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0])
        if err != nil {
            continue // Skip lines with invalid timestamp
        }

        entries = append(entries, LogEntry{Timestamp: timestamp, Message: parts[1]})
    }

    return entries, scanner.Err()
}

func main() {
    e := echo.New()
    e.GET("/parse", func(c echo.Context) error {
        logFilePath := c.QueryParam("file")
        if logFilePath == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Log file path is required")
        }

        logEntries, err := parseLogFile(logFilePath)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to parse log file: %v", err))
        }

        return c.JSON(http.StatusOK, logEntries)
    })

    e.Logger.Fatal(e.Start(":8080"))
}