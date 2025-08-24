// 代码生成时间: 2025-08-24 14:09:32
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/robfig/cron/v3"
    "github.com/labstack/echo/v4"
)

// SchedulerService 结构体，用于封装定时任务的逻辑
type SchedulerService struct {
    cron *cron.Cron
}

// NewSchedulerService 初始化并返回 SchedulerService 结构体的实例
func NewSchedulerService() *SchedulerService {
    return &SchedulerService{
        cron: cron.New(),
    }
}

// Start 启动定时任务调度器
func (s *SchedulerService) Start() {
    _, err := s.cron.AddFunc("*/5 * * * *", func() { s.runTask() })
    if err != nil {
        fmt.Println("Failed to add cron job: ", err)
        return
    }
    s.cron.Start()
    fmt.Println("Cron scheduler started")
}

// Stop 停止定时任务调度器
func (s *SchedulerService) Stop() {
    s.cron.Stop()
    fmt.Println("Cron scheduler stopped")
}

// runTask 定时执行的任务函数
func (s *SchedulerService) runTask() {
    fmt.Println("Running scheduled task...")
    // 在这里添加实际的任务逻辑
}

func main() {
    scheduler := NewSchedulerService()
    defer scheduler.Stop()

    e := echo.New()
    e.GET("/start", func(c echo.Context) error {
        scheduler.Start()
        return c.JSON(200, map[string]string{"message": "Scheduler started"})
    })
    e.GET("/stop", func(c echo.Context) error {
        scheduler.Stop()
        return c.JSON(200, map[string]string{"message": "Scheduler stopped"})
    })

    // 启动Echo服务器
    e.Start(":8080")
}
