// 代码生成时间: 2025-08-07 00:45:09
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/lib/pq"
)

// AppConfig holds application configuration
type AppConfig struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

// User represents a user entity
type User struct {
    ID       int
    Username string
    Password string
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    // Database configuration
    config := AppConfig{
        DBHost:     "localhost",
        DBPort:     "5432",
        DBUser:     "postgres",
        DBPassword: "password",
        DBName:     "mydatabase",
    }
    
    // Initialize the database connection
    db, err := sql.Open("postgres", fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DBHost,
        config.DBPort,
        config.DBUser,
        config.DBPassword,
        config.DBName,
    ))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Ensure the database connection is alive
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    
    // Define route for user creation
    e.POST("/users", createUser)
    
    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}

// createUser handles the creation of new users
func createUser(c echo.Context) error {
    var user User
    if err := c.Bind(&user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }
    
    // Hash the password before storing it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    
    // Use parameterized queries to prevent SQL injection
    stmt := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
    err = db.QueryRow(stmt, user.Username, user.Password).Scan(&user.ID)
    if err != nil {
        return err
    }
    
    return c.JSON(http.StatusOK, user)
}