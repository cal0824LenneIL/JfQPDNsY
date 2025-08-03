// 代码生成时间: 2025-08-03 11:50:17
package main

import (
    "net/http"
    "strings"
    echo "github.com/labstack/echo"
)

// Define a struct to hold form data
type FormData struct {
    Username string `form:"username" json:"username"`
    Email    string `form:"email" json:"email"`
    Age      int    `form:"age" json:"age"`
}

// Validator function for the form data
func validateFormData(data *FormData) error {
    // Validate username
    if strings.TrimSpace(data.Username) == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
    }
    // Validate email
    if strings.TrimSpace(data.Email) == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Email is required")
    }
    // Additional email pattern check
    if !strings.Contains(data.Email, "@") {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
    }
    // Validate age
    if data.Age < 0 || data.Age > 130 {
        return echo.NewHTTPError(http.StatusBadRequest, "Age must be between 0 and 130")
    }
    return nil
}

func main() {
    e := echo.New()
    
    // Define a route for form submission with a POST method
    e.POST("/form", func(c echo.Context) error {
        // Create an instance of FormData
        data := new(FormData)
        
        // Bind the form data to the struct
        if err := c.Bind(data); err != nil {
            return err
        }
        
        // Validate the form data
        if err := validateFormData(data); err != nil {
            return err
        }
        
        // If validation is successful, return a success message
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Form data is valid",
        })
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}