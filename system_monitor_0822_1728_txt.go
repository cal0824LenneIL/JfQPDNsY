// 代码生成时间: 2025-08-22 17:28:20
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "os"
    "runtime"
    "time"
    "strings"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/load"
)

// SystemMetrics contains system performance data
type SystemMetrics struct {
    CpuUsage float64 `json:"cpu_usage"`
    MemoryUsage uint64 `json:"memory_usage"`
    DiskUsage float64 `json:"disk_usage"`
    LoadAvg float64 `json:"load_avg"`
}

// GetSystemMetrics fetches system metrics
func GetSystemMetrics() (SystemMetrics, error) {
    var metrics SystemMetrics
    var err error
    
    // Get CPU usage
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return metrics, err
    }
    metrics.CpuUsage = cpuPercent[0]
    
    // Get memory usage
    vmStat, err := mem.VirtualMemory()
    if err != nil {
        return metrics, err
    }
    metrics.MemoryUsage = vmStat.Used
    
    // Get disk usage
    partitions, err := disk.Partitions(true)
    if err != nil {
        return metrics, err
    }
    for _, partition := range partitions {
        if partition.Mountpoint == "/" {
            u, err := disk.Usage(partition.Mountpoint)
            if err != nil {
                return metrics, err
            }
            metrics.DiskUsage = u.UsedPercent
            break
        }
    }
    
    // Get load average
    loadAvg, err := load.Avg()
    if err != nil {
        return metrics, err
    }
    metrics.LoadAvg = loadAvg.Load1
    
    return metrics, nil
}

func main() {
    e := echo.New()
    
    // Define route for system metrics
    e.GET("/metrics", func(c echo.Context) error {
        metrics, err := GetSystemMetrics()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{
                "error": "Failed to fetch system metrics",
            })
        }
        return c.JSON(http.StatusOK, metrics)
    })
    
    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}