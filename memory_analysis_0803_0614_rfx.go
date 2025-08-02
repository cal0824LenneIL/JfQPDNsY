// 代码生成时间: 2025-08-03 06:14:49
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
    "time"

    "github.com/labstack/echo"
)

// MemoryUsage provides details about the memory usage
type MemoryUsage struct {
    Alloc       uint64 `json:"alloc"`       // bytes allocated and not yet freed
    TotAlloc    uint64 `json:"tot_alloc"`    // total bytes allocated (even if freed)
    Sys         uint64 `json:"sys"`         // bytes obtained from system (sum of XxxSys below)
    Looks       uint64 `json:"looks"`       // number of pointer lookups
    Mallocs     uint64 `json:"mallocs"`     // number of mallocs
    Frees       uint64 `json:"frees"`       // number of frees
    HeapAlloc   uint64 `json:"heap_alloc"`   // bytes allocated in heap
    HeapSys     uint64 `json:"heap_sys"`     // heap system bytes
    HeapIdle    uint64 `json:"heap_idle"`    // heap idle bytes
    HeapInuse   uint64 `json:"heap_inuse"`   // heap in-use bytes
    HeapReleased uint64 `json:"heap_released"` // bytes released to the OS
}

// getMemoryUsage retrieves the current memory usage
func getMemoryUsage() (MemoryUsage, error) {
    m := new(runtime.MemStats)
    runtime.ReadMemStats(m)

    return MemoryUsage{
        Alloc:       m.Alloc,
        TotAlloc:    m.TotalAlloc,
        Sys:         m.Sys,
        Looks:       m.Lookups,
        Mallocs:     m.Mallocs,
        Frees:       m.Frees,
        HeapAlloc:   m.HeapAlloc,
        HeapSys:     m.HeapSys,
        HeapIdle:    m.HeapIdle,
        HeapInuse:   m.HeapInuse,
        HeapReleased: m.HeapReleased,
    }, nil
}

// memoryHandler responds with the current memory usage
func memoryHandler(c echo.Context) error {
    mu, err := getMemoryUsage()
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, mu)
}

func main() {
    e := echo.New()
    // Define a route for the memory usage endpoint
    e.GET("/memory", memoryHandler)

    // Start the server
    if err := e.Start(":8080"); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}
