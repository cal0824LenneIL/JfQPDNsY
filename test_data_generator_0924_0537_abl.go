// 代码生成时间: 2025-09-24 05:37:00
package main

import (
    "crypto/rand"
    "encoding/json"
    "fmt"
    "math/big"
    "net/http"
    "time"

    "github.com/labstack/echo"
)

// TestData 代表测试数据的结构
type TestData struct {
    ID        int64     `json:"id"`
    FirstName string   `json:"first_name"`
    LastName  string   `json:"last_name"`
    Email     string   `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// generateRandomString 生成一个随机字符串
func generateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

// generateTestData 生成测试数据
func generateTestData() TestData {
    return TestData{
        ID:        time.Now().Unix(),
        FirstName: generateRandomString(10),
        LastName:  generateRandomString(10),
        Email:     generateRandomString(10) + "@example.com",
        CreatedAt: time.Now(),
    }
}

// handler 生成并返回测试数据的HTTP处理器
func handler(c echo.Context) error {
    data := generateTestData()
    b, err := json.Marshal(data)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(http.StatusOK, echo.Map{
        "data": string(b),
    })
}

func main() {
    e := echo.New()
    e.GET("/test-data", handler)
    e.Logger.Fatal(e.Start(":8080"))
}
