// 代码生成时间: 2025-09-06 02:11:09
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"
    "strconv"
    "strings"

    "github.com/disintegration/imaging"
    "github.com/labstack/echo"
)

// ImageResizeResponse 定义了图片尺寸调整后的响应结构
type ImageResizeResponse struct {
    Width  int    `json:"width"`
    Height int    `json:"height"`
    Format string `json:"format"`
}

func main() {
    e := echo.New()
    e.POST("/resize", resizeImageHandler)
    e.GET("/images", listImagesHandler)
    e.GET("/image/:filename", serveImageHandler)
    e.Static("/uploads", "uploads")
    e.Logger.Fatal(e.Start(":8080"))
}

// resizeImageHandler 处理图片尺寸调整的POST请求
func resizeImageHandler(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    dst, err := imaging.Open(src, true)
    if err != nil {
        return err
    }
    defer dst.Close()

    widthStr := c.QueryParam("width")
    heightStr := c.QueryParam("height")
    width, err := strconv.Atoi(widthStr)
    if err != nil {
        return err
    }
    height, err := strconv.Atoi(heightStr)
    if err != nil {
        return err
    }

    resized, err := imaging.Resize(dst, width, height, imaging.Lanczos)
    if err != nil {
        return err
    }
    filename := file.Filename + "_resized." + GetImageFormat(dst)
    err = saveImage(resized, filename)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, ImageResizeResponse{
        Width:  resized.Bounds().Dx(),
        Height: resized.Bounds().Dy(),
        Format: GetImageFormat(resized),
    })
}

// listImagesHandler 列出上传目录中的所有图片
func listImagesHandler(c echo.Context) error {
    dir, err := ioutil.ReadDir("uploads")
    if err != nil {
        return err
    }
    var files []string
    for _, d := range dir {
        if d.IsDir() {
            continue
        }
        files = append(files, d.Name())
    }
    sort.Strings(files)
    return c.JSON(http.StatusOK, files)
}

// serveImageHandler 用于提供已上传的图片
func serveImageHandler(c echo.Context) error {
    filename := c.Param("filename")
    file, err := os.Open(filepath.Join("uploads", filename))
    if err != nil {
        return err
    }
    defer file.Close()

    return c.Stream(http.StatusOK, http.DetectContentType(file), file)
}

// saveImage 保存图片到文件系统
func saveImage(img imaging.Image, filename string) error {
    file, err := os.Create(filepath.Join("uploads", filename))
    if err != nil {
        return err
    }
    defer file.Close()

    return imaging.Save(img, file)
}

// GetImageFormat 根据图片内容返回图片格式
func GetImageFormat(img imaging.Image) string {
    return strings.Split(img.Format(), "/")[1]
}
