// 代码生成时间: 2025-09-19 09:57:15
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "encoding/json"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Price   float64 `json:"price"`
    Quantity int    `json:"quantity"`
}

// Cart represents the shopping cart
type Cart struct {
    Items map[int]CartItem
}

// NewCart initializes a new shopping cart
func NewCart() *Cart {
    return &Cart{Items: make(map[int]CartItem)}
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(item CartItem) error {
    if _, exists := c.Items[item.ID]; exists {
        // Increase quantity if item already exists
        c.Items[item.ID].Quantity += item.Quantity
    } else {
        // Add new item to the cart
        c.Items[item.ID] = item
    }
    return nil
}

// RemoveItem removes an item from the cart
func (c *Cart) RemoveItem(itemID int) error {
    if _, exists := c.Items[itemID]; exists {
        delete(c.Items, itemID)
    } else {
        return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
    }
    return nil
}

// UpdateItem updates an item in the cart
func (c *Cart) UpdateItem(itemID int, quantity int) error {
    if cartItem, exists := c.Items[itemID]; exists {
        c.Items[itemID] = CartItem{ID: cartItem.ID, Name: cartItem.Name, Price: cartItem.Price, Quantity: quantity}
    } else {
        return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
    }
    return nil
}

// GetCart returns the current state of the cart
func (c *Cart) GetCart() ([]CartItem, error) {
    var items []CartItem
    for _, item := range c.Items {
        items = append(items, item)
    }
    return items, nil
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    cart := NewCart()
    
    // Add item to cart
    e.POST("/cart/add", func(c echo.Context) error {
        var item CartItem
        if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
            return err
        }
        if err := cart.AddItem(item); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, item)
    })
    
    // Remove item from cart
    e.POST("/cart/remove", func(c echo.Context) error {
        itemID := c.QueryParam("id")
        if _, err := strconv.Atoi(itemID); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid item ID")
        }
        if err := cart.RemoveItem(itemID); err != nil {
            return err
        }
        return c.NoContent(http.StatusOK)
    })
    
    // Update item in cart
    e.POST("/cart/update", func(c echo.Context) error {
        itemIDStr := c.QueryParam("id")
        itemID, err := strconv.Atoi(itemIDStr)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid item ID")
        }
        var quantity int
        if err := json.NewDecoder(c.Request().Body).Decode(&quantity); err != nil {
            return err
        }
        if err := cart.UpdateItem(itemID, quantity); err != nil {
            return err
        }
        return c.NoContent(http.StatusOK)
    })
    
    // Get cart items
    e.GET("/cart", func(c echo.Context) error {
        items, err := cart.GetCart()
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, items)
    })
    
    e.Logger.Fatal(e.Start(":1323"))
}