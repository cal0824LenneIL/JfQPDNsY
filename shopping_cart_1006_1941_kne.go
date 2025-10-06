// 代码生成时间: 2025-10-06 19:41:59
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// ShoppingCart represents a shopping cart with items
type ShoppingCart struct {
    Items map[string]int
}

// CartService handles shopping cart operations
type CartService struct {
    carts map[string]ShoppingCart
}

// NewCartService creates a new instance of CartService
func NewCartService() *CartService {
    return &CartService{
        carts: make(map[string]ShoppingCart),
    }
}

// AddItem adds an item to the shopping cart
func (s *CartService) AddItem(cartID, itemID string) error {
    if _, exists := s.carts[cartID]; !exists {
        s.carts[cartID] = ShoppingCart{Items: make(map[string]int)}
    }
    s.carts[cartID].Items[itemID]++
    return nil
}

// RemoveItem removes an item from the shopping cart
func (s *CartService) RemoveItem(cartID, itemID string) error {
    if _, exists := s.carts[cartID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Cart not found")
    }
    if _, exists := s.carts[cartID].Items[itemID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
    }
    if s.carts[cartID].Items[itemID] == 1 {
        delete(s.carts[cartID].Items, itemID)
    } else {
        s.carts[cartID].Items[itemID]--
    }
    return nil
}

// GetCart retrieves the shopping cart by ID
func (s *CartService) GetCart(cartID string) (*ShoppingCart, error) {
    if cart, exists := s.carts[cartID]; exists {
        return &cart, nil
    } else {
        return nil, echo.NewHTTPError(http.StatusNotFound, "Cart not found")
    }
}

// main function to start the HTTP server
func main() {
    e := echo.New()
    
    // Create a new CartService instance
    cartService := NewCartService()
    
    e.POST("/add_item", func(c echo.Context) error {
        cartID := c.QueryParam("cartID\)
        itemID := c.QueryParam("itemID\)
        err := cartService.AddItem(cartID, itemID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Item added to cart"})
    })
    
    e.POST("/remove_item", func(c echo.Context) error {
        cartID := c.QueryParam("cartID\)
        itemID := c.QueryParam("itemID\)
        err := cartService.RemoveItem(cartID, itemID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Item removed from cart"})
    })
    
    e.GET("/cart/:cartID", func(c echo.Context) error {
        cartID := c.Param("cartID\)
        cart, err := cartService.GetCart(cartID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, cart)
    })
    
    // Start the server
    e.Logger.Fatal(e.Start(":" + "1323"))
}