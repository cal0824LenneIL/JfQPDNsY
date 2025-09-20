// 代码生成时间: 2025-09-21 00:54:47
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"

    _ "github.com/go-sql-driver/mysql"   // 导入MySQL驱动
    "github.com/labstack/echo/v4"
)

// 数据库配置
const dbConfig = "username:password@tcp(127.0.0.1:3306)/dbname"

// DB全局变量
var DB *sql.DB

func init() {
    // 初始化数据库连接
    var err error
    DB, err = sql.Open("mysql", dbConfig)
    if err != nil {
        log.Fatal(err)
    }

    // 设置连接最大生命周期
    DB.SetConnMaxLifetime(time.Hour)
    // 设置最大空闲连接数
    DB.SetMaxIdleConns(10)
    // 设置最大数据库连接数
    DB.SetMaxOpenConns(100)
}

// 查询处理函数
func queryHandler(c echo.Context) error {
    // 获取查询参数
    query := c.QueryParam("query")
    if query == "" {
        return c.JSON(http.StatusBadRequest, "Query parameter is required")
    }

    // 预处理SQL语句，防止SQL注入
    preparedStmt, err := DB.Prepare("SELECT * FROM users WHERE email = ?")
    if err != nil {
        log.Printf("Error preparing statement: %v", err)
        return c.JSON(http.StatusInternalServerError, "Internal server error")
    }
    defer preparedStmt.Close()

    // 绑定参数
    rows, err := preparedStmt.Query(query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return c.JSON(http.StatusInternalServerError, "Internal server error")
    }
    defer rows.Close()

    // 处理查询结果
    var result []string
    for rows.Next() {
        var email string
        if err := rows.Scan(&email); err != nil {
            log.Printf("Error scanning row: %v", err)
            return c.JSON(http.StatusInternalServerError, "Internal server error")
        }
        result = append(result, email)
    }
    if err := rows.Err(); err != nil {
        log.Printf("Error fetching rows: %v", err)
        return c.JSON(http.StatusInternalServerError, "Internal server error")
    }

    // 返回查询结果
    return c.JSON(http.StatusOK, result)
}

func main() {
    e := echo.New()
    // 注册查询处理路由
    e.GET("/query", queryHandler)

    // 启动服务器
    e.Logger.Fatal(e.Start(":8080"))
}
