// 代码生成时间: 2025-08-23 10:51:17
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Quantity    int     `json:"quantity"`
    Price       float64 `json:"price"`
}

// Cart represents a shopping cart
type Cart struct {
    Items  []*CartItem `json:"items"`
}

// NewCart creates a new shopping cart instance
func NewCart() *Cart {
    return &Cart{Items: make([]*CartItem, 0)}
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(item *CartItem) {
    c.Items = append(c.Items, item)
}

// RemoveItem removes an item from the cart by its ID
func (c *Cart) RemoveItem(itemID string) error {
    for i, item := range c.Items {
        if item.ID == itemID {
            c.Items = append(c.Items[:i], c.Items[i+1:]...)
            return nil
        }
    }
    return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
}

// ItemCount returns the total number of items in the cart
func (c *Cart) ItemCount() int {
    return len(c.Items)
}

// TotalPrice calculates the total price of the items in the cart
func (c *Cart) TotalPrice() float64 {
    var total float64
    for _, item := range c.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

// StartServer starts the server with the cart API
func StartServer() *echo.Echo {
    e := echo.New()
    e.POST("/cart", AddToCart)
    e.GET("/cart", GetCart)
    e.DELETE("/cart/:itemID", RemoveFromCart)
    return e
}

// AddToCart adds an item to the cart
func AddToCart(c echo.Context) error {
    item := new(CartItem)
    if err := c.Bind(item); err != nil {
        return err
    }
    cart := NewCart() // In a real application, you'd load the cart from a database or session
    cart.AddItem(item)
    return c.JSON(http.StatusOK, cart)
}

// GetCart retrieves the cart contents
func GetCart(c echo.Context) error {
    cart := NewCart() // In a real application, you'd load the cart from a database or session
    return c.JSON(http.StatusOK, cart)
}

// RemoveFromCart removes an item from the cart
func RemoveFromCart(c echo.Context) error {
    itemID := c.Param("itemID")
    cart := NewCart() // In a real application, you'd load the cart from a database or session
    if err := cart.RemoveItem(itemID); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, cart)
}

func main() {
    e := StartServer()
    e.Start(":8080") // Start server on port 8080
}