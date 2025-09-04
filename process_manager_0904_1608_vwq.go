// 代码生成时间: 2025-09-04 16:08:21
package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strings"
    "time"
    "github.com/labstack/echo"
)

// ProcessManager 结构体用于存储进程信息
type ProcessManager struct {
    ProcessName string
    IsRunning   bool
}

// NewProcessManager 创建一个新的进程管理器实例
func NewProcessManager(processName string) *ProcessManager {
    return &ProcessManager{ProcessName: processName}
}

// StartProcess 启动一个进程
func (pm *ProcessManager) StartProcess() error {
    cmd := exec.Command(pm.ProcessName)
    if err := cmd.Start(); err != nil {
        return fmt.Errorf("failed to start process: %w", err)
    }
    pm.IsRunning = true
    return nil
}

// StopProcess 停止一个进程
func (pm *ProcessManager) StopProcess() error {
    if !pm.IsRunning {
        return fmt.Errorf("process is not running")
    }
    if err := exec.Command("pkill", "-f", pm.ProcessName).Run(); err != nil {
        return fmt.Errorf("failed to stop process: %w", err)
    }
    pm.IsRunning = false
    return nil
}

// CheckProcessStatus 检查进程状态
func (pm *ProcessManager) CheckProcessStatus() bool {
    cmd := exec.Command("pgrep", "-f", pm.ProcessName)
    _, err := cmd.Output()
    return err == nil
}

// ProcessManagerHandler 处理进程管理的HTTP请求
func ProcessManagerHandler(pm *ProcessManager) func(c echo.Context) error {
    return func(c echo.Context) error {
        switch c.Request().Method {
        case http.MethodGet:
            status := pm.CheckProcessStatus()
            return c.JSON(http.StatusOK, map[string]bool{
                "status": status,
            })
        case http.MethodPost:
            if err := pm.StartProcess(); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{
                    "error": err.Error(),
                })
            }
            return c.JSON(http.StatusOK, map[string]string{
                "message": "Process started successfully",
            })
        case http.MethodDelete:
            if err := pm.StopProcess(); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{
                    "error": err.Error(),
                })
            }
            return c.JSON(http.StatusOK, map[string]string{
                "message": "Process stopped successfully",
            })
        default:
            return c.JSON(http.StatusMethodNotAllowed, map[string]string{
                "error": "Method not allowed",
            })
        }
        return nil
    }
}

func main() {
    e := echo.New()
    pm := NewProcessManager("your_process_name")
    
    // 注册进程管理器路由
    e.GET("/process", ProcessManagerHandler(pm))
    e.POST("/process", ProcessManagerHandler(pm))
    e.DELETE("/process", ProcessManagerHandler(pm))
    
    // 启动ECHO服务器
    e.Start(":8080")
}
