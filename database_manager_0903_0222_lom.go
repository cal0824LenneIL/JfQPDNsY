// 代码生成时间: 2025-09-03 02:22:26
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/session"
    "github.com/labstack/gommon/bytes"
)

// DBConfig holds the configuration for the database
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DBManager is responsible for managing the database connection pool
type DBManager struct {
    db *sql.DB
}

// NewDBManager creates a new instance of DBManager with a given configuration
func NewDBManager(cfg DBConfig) (*DBManager, error) {
    // Create a DSN (Data Source Name) string
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DBName,
    )

    // Open a connection to the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set connection pool parameters
    db.SetMaxOpenConns(25) // Maximum number of open connections to the database.
    db.SetMaxIdleConns(25) // Maximum number of connections in the idle connection pool.
    db.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused.

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Return a new DBManager instance
    return &DBManager{db: db}, nil
}

// Close closes the database connection pool
func (m *DBManager) Close() error {
    return m.db.Close()
}

func main() {
    // Database configuration
    dbConfig := DBConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "mydatabase",
    }

    // Create a new DBManager instance
    dbManager, err := NewDBManager(dbConfig)
    if err != nil {
        log.Fatalf("Failed to create DBManager: %s", err)
    }
    defer dbManager.Close()

    // Create an Echo instance
    e := echo.New()

    // Middleware: Session
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

    // Routes
    e.GET("/db", func(c echo.Context) error {
        // Use the database connection
        result, err := dbManager.db.Query("SELECT 1")
        if err != nil {
            log.Println("Error querying database: ", err)
            return err
        }
        defer result.Close()

        // Fetch the first row
        if rows, err := result.Fetch(); err == nil {
            fmt.Println(rows)
        } else {
            log.Println("Error fetching rows: ", err)
        }

        // Return a response
        return c.String(bytes.StatusOK, "Database connection successful")
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
