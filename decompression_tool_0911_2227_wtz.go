// 代码生成时间: 2025-09-11 22:27:55
package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/labstack/echo/v4"
)

// Decompressor 定义一个解压器结构体，包含解压所需的文件名和缓冲区
type Decompressor struct {
    reader *gzip.Reader
    writer io.Writer
}

// NewDecompressor 创建一个新的解压器实例
func NewDecompressor(reader io.Reader, writer io.Writer) *Decompressor {
    var buf bytes.Buffer
    return &Decompressor{
        reader: gzip.NewReader(reader),
        writer: &buf,
    }
}

// Decompress 执行解压操作
func (d *Decompressor) Decompress() error {
    _, err := io.Copy(d.writer, d.reader)
    if err != nil {
        return err
    }
    d.reader.Close()
    return nil
}

// WriteOutput 将解压后的数据写入文件
func (d *Decompressor) WriteOutput(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.Write(d.writer.(*bytes.Buffer).Bytes())
    return err
}

func main() {
    e := echo.New()

    // 定义解压文件的路由
    e.POST("/decompress", func(c echo.Context) error {
        file, err := c.FormFile("file")
        if err != nil {
            return err
        }
        src, err := file.Open()
        if err != nil {
            return err
        }
        defer src.Close()

        // 创建解压器实例
        decompressor := NewDecompressor(src, os.Stdout)
        if err := decompressor.Decompress(); err != nil {
            return err
        }

        // 将解压后的数据写入文件
        outputFilename := fmt.Sprintf("%s_%s", file.Filename, time.Now().Format("20060102_150405"))
        if err := decompressor.WriteOutput(outputFilename); err != nil {
            return err
        }

        return c.JSON(http.StatusOK, echo.Map{
            "message": "File decompressed successfully",
            "output": outputFilename,
        })
    })

    // 启动Echo服务器
    e.Logger.Fatal(e.Start(":8080"))
}
