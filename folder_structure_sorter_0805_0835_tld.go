// 代码生成时间: 2025-08-05 08:35:00
package main

import (
    "bytes"
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// FolderItem represents a file or directory within the folder structure.
type FolderItem struct {
    Name     string    `json:"name"`
    Path     string    `json:"path"`
    IsDir    bool      `json:"is_dir"`
    Size     int64     `json:"size"`
    ModTime  time.Time `json:"mod_time"`
}

// FolderStructureSorter sorts the directory structure into a structured format.
type FolderStructureSorter struct {
    basePath string
}

// NewFolderStructureSorter creates a new FolderStructureSorter instance.
func NewFolderStructureSorter(basePath string) *FolderStructureSorter {
    return &FolderStructureSorter{basePath: basePath}
}

// Sort walks through the directory structure from the base path and sorts the items.
func (f *FolderStructureSorter) Sort() ([]FolderItem, error) {
    var items []FolderItem
    err := filepath.WalkDir(f.basePath, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        info, err := d.Info()
        if err != nil {
            return err
        }
        relativePath, err := filepath.Rel(f.basePath, path)
        if err != nil {
            return err
        }
        items = append(items, FolderItem{
            Name:     d.Name(),
            Path:     filepath.ToSlash(relativePath),
            IsDir:    false,
            Size:     info.Size(),
            ModTime:  info.ModTime(),
        })
        return nil
    })
    if err != nil {
        return nil, err
    }
    // Sort items by path length and then by name.
    sort.SliceStable(items, func(i, j int) bool {
        return items[i].Path < items[j].Path || (items[i].Path == items[j].Path && items[i].Name < items[j].Name)
    })
    return items, nil
}

func main() {
    sorter := NewFolderStructureSorter(".") // Use the current directory as base path.
    items, err := sorter.Sort()
    if err != nil {
        log.Fatalf("Failed to sort folder structure: %v", err)
    }
    // Print out the sorted folder items.
    var buffer bytes.Buffer
    for _, item := range items {
        fmt.Fprintf(&buffer, "{Name: %q, Path: %q, IsDir: %v, Size: %d, ModTime: %v}",
            item.Name, item.Path, item.IsDir, item.Size, item.ModTime)
        buffer.WriteString("
")
    }
    fmt.Println(buffer.String())
}
