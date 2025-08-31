// 代码生成时间: 2025-08-31 17:40:02
Features:
- Process listing
- Process killing
- Health check of the process manager service

This program follows the best practices of Go programming and is designed for maintainability and scalability.
*/

package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os/exec"
    "strings"
    "syscall"

    "github.com/labstack/echo"
)

// ProcessInfo holds information about a process
type ProcessInfo struct {
    PID     int    "json:"pid""
    Name    string "json:"name""
    Created string "json:"created""
}

// ProcessManager is responsible for managing processes
type ProcessManager struct {
    Processes map[int]ProcessInfo
}

// NewProcessManager creates a new instance of ProcessManager
func NewProcessManager() *ProcessManager {
    return &ProcessManager{
        Processes: make(map[int]ProcessInfo),
    }
}

// ListProcesses returns a list of running processes
func (pm *ProcessManager) ListProcesses() ([]ProcessInfo, error) {
    processes, err := pm.fetchProcesses()
    if err != nil {
        return nil, err
