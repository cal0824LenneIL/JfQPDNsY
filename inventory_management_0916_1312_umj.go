// 代码生成时间: 2025-09-16 13:12:43
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/labstack/echo"
)

// InventoryItem represents the structure of an inventory item
type InventoryItem struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Quantity  int    `json:"quantity"`
}

// Database represents a mock database for inventory items
var Database = []InventoryItem{
    {ID: 1, Name: "Widget", Quantity: 100},
    {ID: 2, Name: "Gadget", Quantity: 150},
}

// Handler for adding an inventory item
func addItem(c echo.Context) error {
    var item InventoryItem
    if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid item data").SetInternal(err)
    }
    item.ID = len(Database) + 1 // Assign a new ID
    Database = append(Database, item) // Add to the database
    return c.JSON(http.StatusOK, item)
}

// Handler for getting all inventory items
func getAllItems(c echo.Context) error {
    return c.JSON(http.StatusOK, Database)
}

// Handler for getting an inventory item by ID
func getItemByID(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }
    for _, item := range Database {
        if item.ID == id {
            return c.JSON(http.StatusOK, item)
        }
    }
    return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

// Handler for updating an inventory item
func updateItem(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }
    var item InventoryItem
    if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid item data").SetInternal(err)
    }
    for i, existingItem := range Database {
        if existingItem.ID == id {
            Database[i] = item // Update the item in the database
            return c.JSON(http.StatusOK, item)
        }
    }
    return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

// Handler for deleting an inventory item
func deleteItem(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
    }
    for i, item := range Database {
        if item.ID == id {
            Database = append(Database[:i], Database[i+1:]...) // Remove the item from the database
            return c.NoContent(http.StatusNoContent)
        }
    }
    return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

func main() {
    e := echo.New()
    e.POST("/items", addItem)
    e.GET("/items", getAllItems)
    e.GET("/items/:id", getItemByID)
    e.PUT("/items/:id", updateItem)
    e.DELETE("/items/:id", deleteItem)
    e.Logger.Fatal(e.Start(":8080"))
}
