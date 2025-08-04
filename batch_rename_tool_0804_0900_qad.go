// 代码生成时间: 2025-08-04 09:00:59
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/labstack/echo"
)

// BatchRenameTool 结构体用于封装批量重命名工具的功能
type BatchRenameTool struct {
    BasePath string
}

// NewBatchRenameTool 创建一个新的批量重命名工具实例
func NewBatchRenameTool(basePath string) *BatchRenameTool {
    return &BatchRenameTool{BasePath: basePath}
}

// RenameFiles 实现批量重命名文件的功能，pattern 是旧文件名的模式，
// newBaseName 是新文件名的前缀
func (b *BatchRenameTool) RenameFiles(pattern, newBaseName string) error {
    files, err := filepath.Glob(filepath.Join(b.BasePath, pattern))
    if err != nil {
        return fmt.Errorf("error finding files: %w", err)
    }

    for index, file := range files {
        if _, err := os.Stat(file); os.IsNotExist(err) {
            log.Printf("File %s does not exist, skipping...
", file)
            continue
        }

        // Generate new file name by appending index to the newBaseName
        fileName := fmt.Sprintf("%s_%d%s", newBaseName, index, strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)))
        newFilePath := filepath.Join(b.BasePath, fileName+filepath.Ext(file))

        if err := os.Rename(file, newFilePath); err != nil {
            return fmt.Errorf("error renaming file %s to %s: %w", file, newFilePath, err)
        }

        fmt.Printf("Renamed %s to %s
", file, newFilePath)
    }

    return nil
}

func main() {
    e := echo.New()

    // Define route for renaming files
    e.POST("/rename", func(c echo.Context) error {
        basePath := c.QueryParam("base_path")
        pattern := c.QueryParam("pattern")
        newBaseName := c.QueryParam("new_base_name")

        if basePath == "" || pattern == "" || newBaseName == "" {
            return c.JSON(400, map[string]string{
                "error": "Missing base_path, pattern, or new_base_name parameter",
            })
        }

        tool := NewBatchRenameTool(basePath)
        if err := tool.RenameFiles(pattern, newBaseName); err != nil {
            return c.JSON(500, map[string]string{
                "error": err.Error(),
            })
        }

        return c.JSON(200, map[string]string{
            "message": "Files renamed successfully",
        })
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
