// 代码生成时间: 2025-09-24 15:49:34
package main

import (
    "flag"
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/disintegration/imaging"
    "github.com/labstack/echo/v4"
)

const (
    imageSizeWidth  = 800 // 调整后的图片宽度
    imageSizeHeight = 600 // 调整后的图片高度
)

// ImageResizer 结构体包含图片处理所需的参数
type ImageResizer struct {
    srcPath  string // 源文件夹路径
    dstPath  string // 目标文件夹路径
    width   int    // 目标图片宽度
    height  int    // 目标图片高度
}

// NewImageResizer 初始化ImageResizer
func NewImageResizer(srcPath, dstPath string) *ImageResizer {
    return &ImageResizer{
        srcPath: srcPath,
        dstPath: dstPath,
        width:  imageSizeWidth,
        height: imageSizeHeight,
    }
}

// ResizeAllImages 调整指定文件夹下所有图片尺寸
func (ir *ImageResizer) ResizeAllImages() error {
    err := filepath.WalkDir(ir.srcPath, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil // 忽略子文件夹
        }

        if !strings.HasPrefix(filepath.Ext(d.Name()), ".") {
            return nil // 忽略非图片文件
        }

        img, err := imaging.Open(path)
        if err != nil {
            return err
        }

        resizedImg := imaging.Resize(img, ir.width, ir.height, imaging.Lanczos)
        if err != nil {
            return err
        }

        dstPath := filepath.Join(ir.dstPath, d.Name())
        if err := imaging.Save(resizedImg, dstPath); err != nil {
            return err
        }
        fmt.Printf("Resized image saved: %s
", dstPath)
        return nil
    })
    return err
}

// StartServer 初始化并启动Echo服务器
func StartServer(resizer *ImageResizer) {
    e := echo.New()
    e.POST("/resize", func(c echo.Context) error {
        srcPath := c.FormValue("srcPath")
        dstPath := c.FormValue("dstPath")
        resizer.srcPath = srcPath
        resizer.dstPath = dstPath
        if err := resizer.ResizeAllImages(); err != nil {
            return err
        }
        return c.JSON(200, map[string]string{"message": "Images resized successfully"})
    })

    e.GET("/health", func(c echo.Context) error {
        return c.JSON(200, map[string]string{"status": "ok"})
    })

    e.Logger.Fatal(e.Start(":8080")) // 启动服务器
}

func main() {
    var srcPath, dstPath string
    flag.StringVar(&srcPath, "src", "./src", "Source directory path")
    flag.StringVar(&dstPath, "dst", "./dst", "Destination directory path")
    flag.Parse()

    resizer := NewImageResizer(srcPath, dstPath)
    if err := resizer.ResizeAllImages(); err != nil {
        log.Fatalf("Error resizing images: %v", err)
    }

    // 启动Echo服务器
    StartServer(resizer)
}
