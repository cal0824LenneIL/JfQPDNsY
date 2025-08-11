// 代码生成时间: 2025-08-11 17:10:47
package main

import (
    "echo"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

// Define the structure for handling rename request
type RenameRequest struct {
    SourceFolder string   `json:"source_folder"`
    NewNames     []string `json:"new_names"`
    OldNames     []string `json:"old_names"`
}

// Renamer service to handle the renaming logic
type Renamer struct{}

// NewRenamer creates a new instance of Renamer
func NewRenamer() *Renamer {
    return &Renamer{}
}

// RenameFiles renames files based on the provided request
func (r *Renamer) RenameFiles(req RenameRequest) error {
    for i, oldName := range req.OldNames {
        newPath := filepath.Join(req.SourceFolder, req.NewNames[i])
        oldPath := filepath.Join(req.SourceFolder, oldName)
        if _, err := os.Stat(oldPath); os.IsNotExist(err) {
            return fmt.Errorf("file %s does not exist", oldPath)
        }
        if err := os.Rename(oldPath, newPath); err != nil {
            return fmt.Errorf("failed to rename %s to %s: %w", oldPath, newPath, err)
        }
    }
    return nil
}

// SetupEcho is the entry point for setting up Echo routes
func SetupEcho(e *echo.Echo) {
    e.POST("/rename", renameHandler)
}

// renameHandler handles the HTTP POST request to rename files
func renameHandler(c echo.Context) error {
    req := new(RenameRequest)
    if err := c.Bind(req); err != nil {
        return err
    }
    renamer := NewRenamer()
    if err := renamer.RenameFiles(*req); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, "Files renamed successfully")
}

func main() {
    e := echo.New()
    SetupEcho(e)
    e.Logger.Fatal(e.Start(":8080"))
}