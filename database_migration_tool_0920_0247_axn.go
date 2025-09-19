// 代码生成时间: 2025-09-20 02:47:40
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/session"
    "github.com/go-sql-driver/mysql"
    "github.com/gobuffalo/packr/v2"
    "github.com/jinzhu/gorm"
)

// DatabaseConfig holds configuration for the database connection
type DatabaseConfig struct {
    User     string
    Password string
    Protocol string
    Host     string
    Port     string
    Dbname   string
}

// MigrationTool is the struct that encapsulates the database migration tool functionality
type MigrationTool struct {
    db *gorm.DB
}

// NewMigrationTool creates a new instance of MigrationTool
func NewMigrationTool(cfg DatabaseConfig) (*MigrationTool, error) {
    // Create a DSN (Data Source Name) string for connecting to the database
    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.Dbname)
    
    // Connect to the database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Return a new MigrationTool instance
    return &MigrationTool{db: db}, nil
}

// RunMigrations runs the database migrations
func (mt *MigrationTool) RunMigrations() error {
    // Use packr to load migration files from the /migrations directory
    box := packr.New("migrations", "./migrations")
    
    // Run migrations
    if err := mt.db.Set("gorm:table", "migrations").AutoMigrate(&Migration{}); err != nil {
        return err
    }
    
    // Loop through each file in the box and run the migration
    for _, file := range box.List() {
        migrationData, err := box.FindString(file)
        if err != nil {
            return err
        }
        
        // Run the migration
        if err := mt.db.Exec(migrationData).Error; err != nil {
            return err
        }
    }
    
    return nil
}

// Migration struct represents a database migration
type Migration struct {
    Version string
}

func main() {
    // Define the database configuration
    dbConfig := DatabaseConfig{
        User:     "user",
        Password: "password",
        Protocol: "tcp",
        Host:     "localhost",
        Port:     "3306",
        Dbname:   "dbname",
    }

    // Create a new MigrationTool instance
    mt, err := NewMigrationTool(dbConfig)
    if err != nil {
        log.Fatalf("Failed to create migration tool: %v", err)
    }

    // Run the migrations
    if err := mt.RunMigrations(); err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    fmt.Println("Database migrations completed successfully.")
}
