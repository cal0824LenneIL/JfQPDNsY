// 代码生成时间: 2025-09-07 03:24:33
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/labstack/echo"
)

// DB is a type that wraps *sql.DB
type DB struct{
    *sql.DB
}

// NewDB creates a new database handle
func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

// Close closes the database handle
func (db *DB) Close() error {
    return db.DB.Close()
}

// main is the entry point of the application
func main() {
    e := echo.New()
    // Define your routes
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to the Echo SQL Injection Protection Example")
    })

    // Route to demonstrate SQL injection protection
    e.GET("/users/:id", func(c echo.Context) error {
        idParam := c.Param("id")
        // Use parameterized queries to prevent SQL injection
        query := `SELECT * FROM users WHERE id = ?`
        rows, err := db.Query(query, idParam)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{
                "error": err.Error(),
            })
        }
        defer rows.Close()

        var user struct{
            ID   int
            Name string
        }
        if rows.Next() {
            if err = rows.Scan(&user.ID, &user.Name); err != nil {
                return c.JSON(http.StatusInternalServerError, echo.Map{
                    "error": err.Error(),
                })
            }
            return c.JSON(http.StatusOK, echo.Map{
                "id":   user.ID,
                "name": user.Name,
            })
        }
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "User not found",
        })
    })

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}

// db is the global database handle
var db *DB

func init() {
    var err error
    db, err = NewDB("root:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()
}
