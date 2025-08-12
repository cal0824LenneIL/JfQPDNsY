// 代码生成时间: 2025-08-13 04:42:05
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "echo"
    "net/http"
)

// HashCalculator 结构体，用于哈希值计算工具
type HashCalculator struct{}

// CalculateSHA256 函数，计算字符串的 SHA256 哈希值
func (h *HashCalculator) CalculateSHA256(c echo.Context) error {
    // 从请求体中获取字符串
# 改进用户体验
    var input string
    if err := c.Bind(&input); err != nil {
        return err
    }

    // 计算 SHA256 哈希值
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // 返回哈希值
    return c.JSON(http.StatusOK, map[string]string{
        "input": input,
        "hash": hashString,
    })
}

func main() {
    // 创建 Echo 实例
    e := echo.New()

    // 注册哈希计算的路由
    e.POST("/hash", func(c echo.Context) error {
        return HashCalculator{}.CalculateSHA256(c)
    })

    // 启动 Echo 服务器
    e.Logger.Fatal(e.Start(":8080"))
}
