// 代码生成时间: 2025-09-23 00:33:28
package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "log"

    "github.com/labstack/echo"
)

// InventoryItem represents a single item in the inventory
type InventoryItem struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
    Description string `json:"description"`
}

// Inventory represents the entire inventory
type Inventory struct {
    items map[string]InventoryItem
}

// NewInventory returns a new Inventory instance
func NewInventory() *Inventory {
    return &Inventory{
        items: make(map[string]InventoryItem),
    }
}

// AddItem adds a new item to the inventory
func (i *Inventory) AddItem(item InventoryItem) error {
    if _, exists := i.items[item.ID]; exists {
        return fmt.Errorf("item with id %s already exists", item.ID)
    }
    i.items[item.ID] = item
    return nil
}

// GetItem retrieves an item from the inventory by ID
func (i *Inventory) GetItem(id string) (*InventoryItem, error) {
    item, exists := i.items[id]
    if !exists {
        return nil, fmt.Errorf("item with id %s not found", id)
    }
    return &item, nil
}

// UpdateItem updates an existing item in the inventory
func (i *Inventory) UpdateItem(id string, newItem InventoryItem) error {
    if _, exists := i.items[id]; !exists {
        return fmt.Errorf("item with id %s not found", id)
    }
    newItem.ID = id
    i.items[id] = newItem
    return nil
}

// DeleteItem removes an item from the inventory by ID
func (i *Inventory) DeleteItem(id string) error {
    if _, exists := i.items[id]; !exists {
        return fmt.Errorf("item with id %s not found", id)
    }
    delete(i.items, id)
    return nil
}

// InventoryHandler handles HTTP requests for inventory operations
func InventoryHandler(e *echo.Echo, inventory *Inventory) {
    e.GET("/items/:id", func(c echo.Context) error {
        id := c.Param("id\)
        item, err := inventory.GetItem(id)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, item)
    })

    e.POST("/items", func(c echo.Context) error {
        var item InventoryItem
        if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
            return err
        }
        if err := inventory.AddItem(item); err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, item)
    })

    e.PUT("/items/:id", func(c echo.Context) error {
        id := c.Param("id\)
        var newItem InventoryItem
        if err := json.NewDecoder(c.Request().Body).Decode(&newItem); err != nil {
            return err
        }
        if err := inventory.UpdateItem(id, newItem); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, newItem)
    })

    e.DELETE("/items/:id", func(c echo.Context) error {
        id := c.Param("id\)
        if err := inventory.DeleteItem(id); err != nil {
            return err
        }
        return c.NoContent(http.StatusNoContent)
    })
}

func main() {
    e := echo.New()
    inventory := NewInventory()
    InventoryHandler(e, inventory)
    
    // Start the server
    log.Fatal(e.Start(":8080"))
}
