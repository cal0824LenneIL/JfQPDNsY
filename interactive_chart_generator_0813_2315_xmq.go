// 代码生成时间: 2025-08-13 23:15:08
 * interactive_chart_generator.go
 * This program is an interactive chart generator using GoLang and Echo framework.
 */

package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/wcharczuk/go-chart/v2"
    "os"
)

// ChartGenerator defines the structure for chart generation
type ChartGenerator struct {
    Title    string
    Width    int
    Height   int
    XAxis    []string
    YValues  []float64
}

// NewChartGenerator creates a new ChartGenerator instance
func NewChartGenerator(title string, width, height int, xAxis []string, yValues []float64) *ChartGenerator {
    return &ChartGenerator{
        Title:    title,
        Width:    width,
        Height:   height,
        XAxis:    xAxis,
        YValues:  yValues,
    }
}

// GenerateChart generates a chart and saves it as a PNG file
func (c *ChartGenerator) GenerateChart(filePath string) error {
    // Create a new chart
    barChart := chart.BarChart{
        Title: c.Title,
        Width: c.Width,
        Height: c.Height,
        XAxis: chart.XAxis{Name: "X Axis"},
        YAxis: chart.YAxis{Name: "Y Axis"},
    }

    // Add series to the chart
    for i, value := range c.YValues {
        barChart.AddSeriesName(fmt.Sprintf("Product %d", i)).
            AddPoint(c.XAxis[i], value)
    }

    // Render the chart
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    return barChart.Render(chart.PNG, file)
}

// StartServer starts the interactive chart generator server
func StartServer() *echo.Echo {
    e := echo.New()

    // Health check endpoint
    e.GET("/health", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "status": "ok",
        })
    })

    // Endpoint to generate and download a chart
    e.POST("/chart", func(c echo.Context) error {
        title := c.QueryParam("title")
        width, _ := strconv.Atoi(c.QueryParam("width"))
        height, _ := strconv.Atoi(c.QueryParam("height"))
        xAxis := c.QueryParam("xAxis")
        yValuesStr := c.QueryParam("yValues")
        yValues := make([]float64, 0)

        // Parse Y values from string to float64
        if err := json.Unmarshal([]byte(yValuesStr), &yValues); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid Y values format",
            })
        }

        // Create a new chart generator
        chartGen := NewChartGenerator(title, width, height, strings.Split(xAxis, ","), yValues)

        // Generate the chart
        filePath := fmt.Sprintf("./%s.png", title)
        if err := chartGen.GenerateChart(filePath); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to generate chart",
            })
        }

        // Return the chart file path as a response
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Chart generated successfully",
            "filePath": filePath,
        })
    })

    return e
}

func main() {
    e := StartServer()
    e.Logger.Fatal(e.Start(":8080"))
}
