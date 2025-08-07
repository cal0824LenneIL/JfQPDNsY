// 代码生成时间: 2025-08-08 00:16:40
package main

import (
    "encoding/json"
    "net/http"
    "github.com/labstack/echo"
)

// ApiResponse 定义API响应的结构
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
}

// ResponseFormatter 是一个函数，用于格式化响应
func ResponseFormatter(c echo.Context, message string, data interface{}) error {
    response := ApiResponse{
        Success: true,
        Message: message,
        Data:    data,
    }
    return c.JSON(http.StatusOK, response)
}

// ErrorResponseFormatter 是一个函数，用于格式化错误响应
func ErrorResponseFormatter(c echo.Context, code int, message string) error {
    response := ApiResponse{
        Success: false,
        Message: message,
        Data:    nil,
    }
    return c.JSON(code, response)
}

// main 函数初始化Echo实例并设置路由
func main() {
    e := echo.New()
    
    // 定义一个示例路由，返回格式化的响应
    e.GET("/example", func(c echo.Context) error {
        // 模拟数据
        data := map[string]string{"key": "value"}
        return ResponseFormatter(c, "Success", data)
    })
    
    // 定义一个示例路由，返回格式化的错误响应
    e.GET("/error", func(c echo.Context) error {
        return ErrorResponseFormatter(c, http.StatusInternalServerError, "Internal Server Error")
    })

    // 启动服务器
    e.Logger.Fatal(e.Start(":8080"))
}