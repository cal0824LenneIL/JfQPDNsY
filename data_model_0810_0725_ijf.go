// 代码生成时间: 2025-08-10 07:25:23
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

// User represents a user data model.
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"` // Note: In a real-world scenario, you would not store passwords in plain text.
}

// UserRepository interface defines the methods for user-related operations.
type UserRepository interface {
    Create(user *User) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
    Update(user *User) error
    Delete(id int) error
}

// InMemoryUserRepository implements the UserRepository interface using an in-memory store.
type InMemoryUserRepository struct {
    users map[int]*User
}

// NewInMemoryUserRepository creates a new in-memory user repository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: make(map[int]*User),
    }
}

// Create adds a new user to the repository.
func (repo *InMemoryUserRepository) Create(user *User) error {
    if _, exists := repo.users[user.ID]; exists {
        return echo.NewHTTPError(http.StatusBadRequest, "User already exists")
    }
    repo.users[user.ID] = user
    return nil
}

// FindByID retrieves a user by their ID.
func (repo *InMemoryUserRepository) FindByID(id int) (*User, error) {
    user, exists := repo.users[id]
    if !exists {
        return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    return user, nil
}

// FindByEmail retrieves a user by their email address.
func (repo *InMemoryUserRepository) FindByEmail(email string) (*User, error) {
    for _, user := range repo.users {
        if user.Email == email {
            return user, nil
        }
    }
    return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
}

// Update modifies an existing user in the repository.
func (repo *InMemoryUserRepository) Update(user *User) error {
    if _, exists := repo.users[user.ID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    repo.users[user.ID] = user
    return nil
}

// Delete removes a user from the repository by their ID.
func (repo *InMemoryUserRepository) Delete(id int) error {
    if _, exists := repo.users[id]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    delete(repo.users, id)
    return nil
}

func main() {
    e := echo.New()
    repo := NewInMemoryUserRepository()

    // User routes
    e.POST("/users", func(c echo.Context) error {
        user := new(User)
        if err := c.Bind(user); err != nil {
            return err
        }
        if err := repo.Create(user); err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, user)
    })

    e.GET("/users/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
        }
        user, err := repo.FindByID(id)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, user)
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
