// 代码生成时间: 2025-08-25 12:22:55
package main

import (
    "html/template"
    "log"
    "net/http"
    "strings"
    "github.com/labstack/echo"
)

// indexPageTemplate is the HTML template for the index page.
// It is defined as a variable for easy modification and testing.
var indexPageTemplate = `{{define "index"}}<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Responsive Layout</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
        .container {
            width: 80%;
            margin: auto;
            overflow: hidden;
        }
        @media (max-width: 768px) {
            .container {
                width: 95%;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to the Responsive Layout</h1>
        <p>This page is designed to respond to different screen sizes.</p>
    </div>
</body>
</html>{{end}}`

// renderTemplate renders the template with the given data.
func renderTemplate(t *template.Template, name string, data map[string]interface{}, c echo.Context) error {
    if err := t.ExecuteTemplate(c.Response().Writer, name, data); err != nil {
        return err
    }
    return nil
}

func main() {
    // Create a new Echo instance.
    e := echo.New()

    // Define the index page handler.
    e.GET("/", func(c echo.Context) error {
        // Create a map to pass data to the template.
        data := make(map[string]interface{})

        // Render the index page template with the data.
        return renderTemplate(template.Must(template.New("index").Parse(indexPageTemplate)), "index", data, c)
    })

    // Start the Echo server.
    log.Printf("Server is running at http://localhost:%d", 1323)
    if err := e.Start(":" + "1323"); err != nil && strings.Contains(err.Error(), "bind:") {
        log.Fatalf("Could not start server: %s", err)
    }
}