// 代码生成时间: 2025-09-29 00:00:33
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    "github.com/labstack/echo"
)

// LogEntry represents a structured log entry
type LogEntry struct {
    Timestamp string
    Level     string
    Message   string
}

// parseLogLine parses a log line and returns a LogEntry if it's valid
func parseLogLine(line string) (*LogEntry, error) {
    parts := strings.Fields(line)
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log line: %s", line)
    }

    timestamp := parts[0] + " " + parts[1]
    level := parts[2]
    message := strings.Join(parts[3:], " ")

    return &LogEntry{
        Timestamp: timestamp,
        Level:     level,
        Message:   message,
    }, nil
}

func main() {
    e := echo.New()

    // Route for parsing log files
    e.GET("/log/parse", func(c echo.Context) error {
        file, err := c.FormFile("file")
        if err != nil {
            return err
        }
        src, err := file.Open()
        if err != nil {
            return err
        }
        defer src.Close()

        var entries []LogEntry
        scanner := bufio.NewScanner(src)
        for scanner.Scan() {
            line := scanner.Text()
            entry, err := parseLogLine(line)
            if err != nil {
                log.Printf("Error parsing log line: %s
", err)
                continue
            }
            entries = append(entries, *entry)
        }
        if err := scanner.Err(); err != nil {
            return err
        }

        return c.JSON(200, entries)
    })

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
