// 代码生成时间: 2025-10-05 05:02:57
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/fsnotify/fsnotify"
)

// FileWatcher 结构体，用于存储watcher和监控目录
type FileWatcher struct {
    watcher *fsnotify.Watcher
    path    string
}

// NewFileWatcher 创建一个新的文件监控器
func NewFileWatcher(path string) (*FileWatcher, error) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }
    err = watcher.Add(path)
    if err != nil {
        watcher.Close()
        return nil, err
    }
    return &FileWatcher{watcher: watcher, path: path}, nil
}

// Watch 文件监控逻辑
func (fw *FileWatcher) Watch() {
    for {
        select {
        case event, ok := <-fw.watcher.Events:
            if !ok {
                // 监听器关闭
                fmt.Println("Watcher closed")
                return
            }
            // 检查文件是否被修改
            if event.Op&fsnotify.Write == fsnotify.Write {
                fmt.Printf("File %s was modified", event.Name)
            }
        case err, ok := <-fw.watcher.Errors:
            if !ok {
                // 监听器关闭
                fmt.Println("Watcher closed")
                return
            }
            // 错误处理
            log.Printf("Watcher error: %s", err)
        }
    }
}

// StartEchoServer 启动Echo服务器
func StartEchoServer() *echo.Echo {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    return e
}

// main 函数，程序入口
func main() {
    // 创建Echo服务器
    e := StartEchoServer()

    // 启动Echo服务器
    go func() {
        if err := e.Start(":8080"); err != nil && err != echo.ErrServerClosed {
            log.Fatalf("shutting down server: %v", err)
        }
    }()

    // 创建文件监控器
    watcher, err := NewFileWatcher("./files")
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.watcher.Close()

    // 启动文件监控
    go func() {
        watcher.Watch()
    }()

    // 等待服务器关闭
    <-e.Context().Done()
}
