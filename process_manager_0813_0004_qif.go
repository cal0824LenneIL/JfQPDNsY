// 代码生成时间: 2025-08-13 00:04:46
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "os/exec"
    "strings"

    "github.com/labstack/echo"
)

// Process represents a system process.
type Process struct {
    Name    string `json:"name"`
    Pid     int    `json:"pid"`
    Running bool   `json:"running"`
}

// ProcessManager handles process operations.
type ProcessManager struct {
    processes map[string]*Process
}

// NewProcessManager creates a new process manager.
func NewProcessManager() *ProcessManager {
    return &ProcessManager{
        processes: make(map[string]*Process),
    }
}

// StartProcess starts a new process with the given command.
func (pm *ProcessManager) StartProcess(name string, command string) (*Process, error) {
    // Split the command into name and args.
    parts := strings.Fields(command)
    if len(parts) == 0 {
        return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid command")
    }
    
    cmd := exec.Command(parts[0], parts[1:]...)
    if err := cmd.Start(); err != nil {
        return nil, err
    }
    
    // Add the process to the manager.
    pm.processes[name] = &Process{
        Name:    name,
        Pid:     cmd.Process.Pid,
        Running: true,
    }
    
    return pm.processes[name], nil
}

// StopProcess stops a process by its name.
func (pm *ProcessManager) StopProcess(name string) error {
    proc, exists := pm.processes[name]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Process not found")
    }
    
    if proc.Running {
        if err := exec.Command("kill", "-9", strconv.Itoa(proc.Pid)).Run(); err != nil {
            return err
        }
        proc.Running = false
    }
    return nil
}

// ListProcesses returns a list of all processes.
func (pm *ProcessManager) ListProcesses() []*Process {
    var procs []*Process
    for _, proc := range pm.processes {
        procs = append(procs, proc)
    }
    return procs
}

// ProcessManagerRoutes sets up the routes for the process manager.
func ProcessManagerRoutes(e *echo.Echo, pm *ProcessManager) {
    e.POST("/process/start", func(c echo.Context) error {
        var req struct {
            Name    string `json:"name"`
            Command string `json:"command"`
        }
        if err := json.NewDecoder(c.Request()).Decode(&req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
        }
        
        proc, err := pm.StartProcess(req.Name, req.Command)
        if err != nil {
            return err
        }
        
        return c.JSON(http.StatusOK, proc)
    })

    e.POST("/process/stop", func(c echo.Context) error {
        name := c.QueryParam("name")
        if err := pm.StopProcess(name); err != nil {
            return err
        }
        return c.NoContent(http.StatusOK)
    })

    e.GET("/processes", func(c echo.Context) error {
        procs := pm.ListProcesses()
        return c.JSON(http.StatusOK, procs)
    })
}

func main() {
    e := echo.New()
    pm := NewProcessManager()
    ProcessManagerRoutes(e, pm)
    e.Logger.Fatal(e.Start(":8080"))
}