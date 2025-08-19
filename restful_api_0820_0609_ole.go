// 代码生成时间: 2025-08-20 06:09:46
package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)

// 定义一个用户模型
type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// 初始化Echo实例
func main() {
    e := echo.New()
    
    // 定义路由
    e.GET("/users", getAllUsers)
    e.POST("/users", createUser)
    
    // 启动服务
    e.Logger.Fatal(e.Start(":8080"))
}

// getAllUsers 获取所有用户
func getAllUsers(c echo.Context) error {
    // 模拟数据库查询
    users := []User{
        {ID: 1, Name: "John Doe", Email: "john@example.com"},
        {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
    }
    
    return c.JSON(http.StatusOK, users) // 返回JSON响应
}

// createUser 创建新用户
func createUser(c echo.Context) error {
    // 从请求体中解析用户数据
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err // 错误处理
    }
    
    // 模拟数据库插入
    fmt.Println("Creating user: ", u)
    
    // 返回成功响应
    return c.JSON(http.StatusCreated, u)
}
