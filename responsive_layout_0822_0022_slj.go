// 代码生成时间: 2025-08-22 00:22:57
package main

import (
    "net/http"
    "strings"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// 定义一个响应结构体
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func main() {
    e := echo.New()

    // 使用中间件来处理跨域请求
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
    }))

    // 路由定义
    e.GET("/", func(c echo.Context) error {
        // 获取用户代理信息
        userAgent := c.Request().UserAgent()
        // 根据用户代理信息判断设备类型
        deviceType := getDeviceType(userAgent)
        // 返回响应式布局内容
        return c.JSON(http.StatusOK, Response{
            Status:  "success",
            Message: "Responsive layout for: " + deviceType,
        })
    })

    // 启动服务
    e.Logger.Fatal(e.Start(":8080"))
}

// getDeviceType 函数用于根据用户代理信息判断设备类型
func getDeviceType(userAgent string) string {
    if strings.Contains(userAgent, "Mobile") || strings.Contains(userAgent, "Android") || strings.Contains(userAgent, "webOS") ||
        strings.Contains(userAgent, "iPhone") || strings.Contains(userAgent, "iPod") || strings.Contains(userAgent, "BlackBerry") ||
        strings.Contains(userAgent, "Windows Phone") {
        return "mobile"
    } else if strings.Contains(userAgent, "iPad") || strings.Contains(userAgent, "Android Pad") || strings.Contains(userAgent, "tablet") ||
        strings.Contains(userAgent, "Nexus 7") || strings.Contains(userAgent, "Nexus 9") {
        return "tablet"
    }
    return "desktop"
}
