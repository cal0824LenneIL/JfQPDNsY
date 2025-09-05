// 代码生成时间: 2025-09-05 17:27:36
package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "github.com/labstack/echo"
)

// Item represents a simple item with an ID and Name.
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: 1, Name: "Item 1"},
    {ID: 2, Name: "Item 2"},
    {ID: 3, Name: "Item 3"},
}

// itemID generates an ID for a new item.
func itemID() int {
    id := items[len(items)-1].ID + 1
    return id
}

// getItems returns all items.
func getItems(c echo.Context) error {
    return c.JSON(http.StatusOK, items)
}

// getItem returns an item by its ID.
func getItem(c echo.Context) error {
    id := c.Param("id")
    item := getItemByID(id)
    if item == nil {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "Item not found",
        })
    }
    return c.JSON(http.StatusOK, item)
}

// postItem creates a new item.
func postItem(c echo.Context) error {
    item := new(Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }
    items = append(items, *item)
    return c.JSON(http.StatusCreated, item)
}

// putItem updates an existing item.
func putItem(c echo.Context) error {
    id := c.Param("id")
    item := getItemByID(id)
    if item == nil {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "Item not found",
        })
    }
    itemToUpdate := new(Item)
    if err := c.Bind(itemToUpdate); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }
    item.Name = itemToUpdate.Name
    return c.JSON(http.StatusOK, item)
}

// deleteItem removes an item by its ID.
func deleteItem(c echo.Context) error {
    id := c.Param("id")
    item := getItemByID(id)
    if item == nil {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "Item not found",
        })
    }
    items = removeItemByID(id)
    return c.NoContent(http.StatusNoContent)
}

// getItemByID retrieves an item by its ID.
func getItemByID(id string) *Item {
    for i := range items {
        if items[i].ID == atoi(id) {
            return &items[i]
        }
    }
    return nil
}

// atoi converts string to int.
func atoi(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return i
}

// removeItemByID removes an item by its ID from the slice.
func removeItemByID(id string) []Item {
    index := -1
    for i := range items {
        if items[i].ID == atoi(id) {
            index = i
            break
        }
    }
    if index == -1 {
        return items
    }
    items = append(items[:index], items[index+1:]...)
    return items
}

func main() {
    e := echo.New()
    
    // Define routes.
    e.GET("/items", getItems)
    e.GET("/items/:id", getItem)
    e.POST("/items", postItem)
    e.PUT("/items/:id", putItem)
    e.DELETE("/items/:id", deleteItem)
    
    // Start server.
    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT") + ""))
}
