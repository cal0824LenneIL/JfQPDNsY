// 代码生成时间: 2025-10-01 04:12:34
package main

import (
    "context"
    "net/http"
    "fmt"
    "log"
    "os"
    "github.com/labstack/echo/v4"
)

// IoTDevice represents a device in the agricultural IoT network
type IoTDevice struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    Temperature float64 `json:"temperature"`
    Humidity    float64 `json:"humidity"`
}

// deviceStore is a simple in-memory store for devices
var deviceStore = make(map[string]IoTDevice)

// DeviceHandler handles incoming requests for device data
func DeviceHandler(c echo.Context) error {
    deviceID := c.Param("id")
    device, exists := deviceStore[deviceID]
    if !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Device not found")
    }
    return c.JSON(http.StatusOK, device)
}

// AddDeviceHandler adds a new device to the store
func AddDeviceHandler(c echo.Context) error {
    var device IoTDevice
    if err := c.Bind(&device); err != nil {
        return err
    }
    if _, exists := deviceStore[device.ID]; exists {
        return echo.NewHTTPError(http.StatusConflict, "Device already exists")
    }
    deviceStore[device.ID] = device
    return c.JSON(http.StatusCreated, device)
}

// UpdateDeviceHandler updates an existing device's data
func UpdateDeviceHandler(c echo.Context) error {
    var device IoTDevice
    deviceID := c.Param("id")
    if err := c.Bind(&device); err != nil {
        return err
    }
    if _, exists := deviceStore[deviceID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Device not found")
    }
    deviceStore[deviceID] = device
    return c.JSON(http.StatusOK, device)
}

func main() {
    e := echo.New()
    e.GET("/devices/:id", DeviceHandler)
    e.POST("/devices", AddDeviceHandler)
    e.PUT("/devices/:id", UpdateDeviceHandler)

    // Start the Echo server
    log.Printf("Starting agriculture IoT server on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
