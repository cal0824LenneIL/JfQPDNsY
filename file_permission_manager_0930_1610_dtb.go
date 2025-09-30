// 代码生成时间: 2025-09-30 16:10:54
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"

    "github.com/labstack/echo"
)

// FilePermissionManager 结构体用于封装文件权限管理相关的方法
type FilePermissionManager struct {
    // 这里可以添加更多字段，比如配置信息等
}

// NewFilePermissionManager 创建并返回一个 FilePermissionManager 实例
func NewFilePermissionManager() *FilePermissionManager {
    return &FilePermissionManager{}
}

// ChangeFilePermission 方法允许改变文件的权限
func (m *FilePermissionManager) ChangeFilePermission(file string, permission os.FileMode) error {
    // 使用 os.Chmod 更改文件权限
    if err := os.Chmod(file, permission); err != nil {
        return fmt.Errorf("error changing file permission: %w", err)
    }
    return nil
}

// ChangeFileOwner 方法允许改变文件的所有者和组
func (m *FilePermissionManager) ChangeFileOwner(file string, uid, gid uint32) error {
    // 使用 os.Chown 更改文件所有者和组
    if err := os.Chown(file, int(uid), int(gid)); err != nil {
        return fmt.Errorf("error changing file owner and group: %w", err)
    }
    return nil
}

// FilePermissionHandler 是处理文件权限变更的 HTTP 处理器
func FilePermissionHandler(c echo.Context) error {
    file := c.QueryParam("file")
    permissionStr := c.QueryParam("permission")
    uidStr := c.QueryParam("uid")
    gidStr := c.QueryParam("gid")

    // 将权限字符串转换为 FileMode
    permission, err := strconv.ParseUint(permissionStr, 8, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid permission format. Use octal (e.g., 0644)."
        })
    }

    // 将 UID 和 GID 字符串转换为 uint32
    uid, err := strconv.ParseUint(uidStr, 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid UID format."
        })
    }
    gid, err := strconv.ParseUint(gidStr, 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "Invalid GID format."
        })
    }

    manager := NewFilePermissionManager()

    // 更改文件权限
    if err := manager.ChangeFilePermission(file, os.FileMode(permission)); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": err.Error(),
        })
    }

    // 更改文件所有者和组
    if err := manager.ChangeFileOwner(file, uint32(uid), uint32(gid)); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "File permissions and ownership updated successfully.",
    })
}

func main() {
    e := echo.New()
    e.GET("/change-permission", FilePermissionHandler)

    // 启动 Echo 服务器
    e.Start(":8080")
}
