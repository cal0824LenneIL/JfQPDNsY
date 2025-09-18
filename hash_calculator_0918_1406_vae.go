// 代码生成时间: 2025-09-18 14:06:01
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "echo"
    "net/http"
    "strings"
)

// HashCalculator 结构体用于保存 Echo 实例
type HashCalculator struct {
    e *echo.Echo
}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator() *HashCalculator {
    e := echo.New()
    return &HashCalculator{e: e}
}

// Start 启动服务
func (h *HashCalculator) Start(port string) {
    h.e.Start(port)
}

// SHA256Handler 处理 SHA-256 哈希请求
func (h *HashCalculator) SHA256Handler() echo.HandlerFunc {
    return func(c echo.Context) error {
        // 获取要哈希的字符串
        input := c.QueryParam("input")
        if input == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Input parameter is required")
        }

        // 计算 SHA-256 哈希值
        hash := sha256.Sum256([]byte(input))
        hexHash := hex.EncodeToString(hash[:])

        // 返回哈希值
        return c.JSON(http.StatusOK, map[string]string{
            "hash": hexHash,
        })
    }
}

func main() {
    // 创建 HashCalculator 实例
    calculator := NewHashCalculator()

    // 添加路由
    calculator.e.GET("/hash", calculator.SHA256Handler())

    // 启动服务
    calculator.Start(":8080")
}
