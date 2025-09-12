// 代码生成时间: 2025-09-12 19:33:53
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-co-op/gocron"
    "github.com/labstack/echo"
)

// Scheduler defines the structure for the task scheduler
type Scheduler struct {
    jobScheduler *gocron.Scheduler
}

// NewScheduler creates a new instance of Scheduler
func NewScheduler() *Scheduler {
    s := gocron.NewScheduler(time.UTC)
    return &Scheduler{jobScheduler: s}
}

// ScheduleJob adds a job to the scheduler
func (s *Scheduler) ScheduleJob(spec string, cmd func()) error {
    _, err := s.jobScheduler.Crontab(spec, cmd)
    if err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() error {
    s.jobScheduler.StartBlocking()
    return nil
}

// EchoServer starts an Echo server with the scheduler
func EchoServer() *echo.Echo {
    e := echo.New()

    // Define routes
    e.GET("/", func(c echo.Context) error {
        return c.JSON(200, map[string]string{
            "message": "Hello, World!",
        })
    })

    return e
}

func main() {
    scheduler := NewScheduler()

    // Schedule a job that runs every minute
    err := scheduler.ScheduleJob("* * * * *", func() {
        fmt.Println("Running scheduled job...")
    })
    if err != nil {
        log.Fatal(err)
    }

    // Start the scheduler
    go func() {
        err := scheduler.Start()
        if err != nil {
            log.Fatal(err)
        }
    }()

    // Start the Echo server
    server := EchoServer()
    
    // Handle shutdown signals
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        log.Println("Shutting down...")
        server.Shutdown(context.Background())
        os.Exit(0)
    }()

    // Start the server
    log.Fatal(server.Start(":8080"))
}
