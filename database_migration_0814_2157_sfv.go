// 代码生成时间: 2025-08-14 21:57:37
package main

import (
    "database/sql"
    "fmt"
    "os"
    "log"
    "gopkg.in/gormigrate.v2"
    "github.com/jinzhu/gorm"
    \_ "github.com/jinzhu/gorm/dialects/sqlite" // 导入SQLite驱动
)

// MigrationStep 定义了一个数据库迁移步骤
type MigrationStep struct {
    Version  int    
    Name     string 
    Up       string // SQL语句
    Down     string // SQL语句
}

func main() {
    // 数据库配置
    dbConfig := "file:mydb.sqlite?cache=shared&mode=memory"
    db, err := sql.Open("sqlite", dbConfig)
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()

    // 迁移配置
    m := gormigrate.Gormigrate{
        Db:      db,
        Dialect: "sqlite",
    }

    // 定义迁移步骤
    migrations := []MigrationStep{
        {
            Version: 1,
            Name:    "init",
            Up: `CREATE TABLE IF NOT EXISTS users (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                email TEXT NOT NULL
            );`,
            Down: `DROP TABLE IF EXISTS users;`,
        },
        // 可以添加更多的迁移步骤
    }

    // 应用迁移
    err = m.Migrate(migrations)
    if err != nil {
        log.Fatal("Migration failed: ", err)
    }
    fmt.Println("Migration complete.")
}
