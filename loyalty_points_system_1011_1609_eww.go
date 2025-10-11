// 代码生成时间: 2025-10-11 16:09:52
 * loyalty_points_system.go
 * 
 * This program is a simple loyalty points system using the Go language
 * and the Echo web framework. It handles the basic operations of a loyalty
 * points system, such as adding points to a user's account and redeeming points.
 */

package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

// Member represents a member of the loyalty program.
type Member struct {
    ID       int    `json:"id"`
    Points   int    `json:"points"`
}

// PointsService manages the operations related to points.
type PointsService struct {
    // A mock database for demonstration purposes.
    members map[int]Member
}

// NewPointsService creates a new PointsService instance.
func NewPointsService() *PointsService {
    return &PointsService{
        members: make(map[int]Member),
    }
}

// AddPoints adds points to a member's account.
func (s *PointsService) AddPoints(memberID int, points int) error {
    if points < 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Points cannot be negative.")
    }
    if member, exists := s.members[memberID]; exists {
        s.members[memberID] = Member{ID: memberID, Points: member.Points + points}
    } else {
        s.members[memberID] = Member{ID: memberID, Points: points}
    }
    return nil
}

// RedeemPoints redeems points for a member.
func (s *PointsService) RedeemPoints(memberID int, points int) error {
    if points < 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Points cannot be negative.")
    }
    if member, exists := s.members[memberID]; exists {
        if member.Points >= points {
            s.members[memberID] = Member{ID: memberID, Points: member.Points - points}
            return nil
        }
        return echo.NewHTTPError(http.StatusBadRequest, "Insufficient points.")
    }
    return echo.NewHTTPError(http.StatusNotFound, "Member not found.")
}

// MemberHandler handles HTTP requests for member points operations.
func MemberHandler(service *PointsService) echo.HandlerFunc {
    return func(c echo.Context) error {
        memberID, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid member ID.")
        }
        var points int
        if c.Request().Method == http.MethodPost {
            if err := json.NewDecoder(c.Request()).Decode(&points); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, "Invalid points format.")
            }
            if err := service.AddPoints(memberID, points); err != nil {
                return err
            }
            return c.JSON(http.StatusOK, map[string]string{"message": "Points added successfully."})
        } else if c.Request().Method == http.MethodDelete {
            if err := json.NewDecoder(c.Request()).Decode(&points); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, "Invalid points format.")
            }
            if err := service.RedeemPoints(memberID, points); err != nil {
                return err
            }
            return c.JSON(http.StatusOK, map[string]string{"message": "Points redeemed successfully."})
        }
        return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed.")
    }
}

func main() {
    e := echo.New()
    service := NewPointsService()
    
    // Define the route for adding/redeeming points to a member's account.
    e.POST("/members/:id/add-points", MemberHandler(service))
    e.DELETE("/members/:id/redeem-points", MemberHandler(service))

    // Start the server.
    e.Logger.Fatal(e.Start(":8080"))
}