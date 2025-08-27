// 代码生成时间: 2025-08-27 23:42:52
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "encoding/json"
    "log"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
    Price       float64 `json:"price"`
}

// Cart represents the shopping cart with a list of items
type Cart struct {
    Items []CartItem `json:"items"`
}

// AddItem adds an item to the shopping cart
func (c *Cart) AddItem(item CartItem) {
    c.Items = append(c.Items, item)
}

// GetCart returns the shopping cart items
func GetCart() *Cart {
    return &Cart{}
}

func main() {
    e := echo.New()

    // POST endpoint to add an item to the cart
    e.POST("/cart", func(c echo.Context) error {
        var item CartItem
        if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
            return err
        }
        cart := GetCart()
        cart.AddItem(item)
        return c.JSON(http.StatusOK, cart)
    })

    // GET endpoint to retrieve the cart contents
    e.GET("/cart", func(c echo.Context) error {
        cart := GetCart()
        return c.JSON(http.StatusOK, cart)
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
