// 代码生成时间: 2025-09-15 20:28:44
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/labstack/echo"
)

// PasswordManager 结构体包含AES加密器
type PasswordManager struct {
    cipher cipher.Block
}

// NewPasswordManager 初始化并返回一个PasswordManager实例
func NewPasswordManager(key []byte) (*PasswordManager, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    return &PasswordManager{cipher: block}, nil
}

// Encrypt 加密密码
func (pm *PasswordManager) Encrypt(plaintext string) (string, error) {
    gcm, err := cipher.NewGCM(pm.cipher)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密密码
func (pm *PasswordManager) Decrypt(ciphertext string) (string, error) {
    decoded, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(pm.cipher)
    if err != nil {
        return "", err
    }
    nonceSize := gcm.NonceSize()
    if len(decoded) < nonceSize {
        return "", fmt.Errorf("ciphertext too short")
    }
    nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}

//handler 处理加密和解密的HTTP请求
func handler(pm *PasswordManager) echo.HandlerFunc {
    return func(c echo.Context) error {
        action := c.QueryParam("action")
        password := c.QueryParam("password")
        if password == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "password is required")
        }
        var result string
        var err error
        switch action {
        case "encrypt":
            result, err = pm.Encrypt(password)
        case "decrypt":
            result, err = pm.Decrypt(password)
        default:
            return echo.NewHTTPError(http.StatusBadRequest, "invalid action")
        }
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }
        return c.JSON(http.StatusOK, map[string]string{
            "result": result,
        })
    }
}

func main() {
    key := []byte("your-32-byte-long-key")
    pm, err := NewPasswordManager(key)
    if err != nil {
        log.Fatal(err)
    }
    e := echo.New()
    e.GET("/password", handler(pm))
    e.Logger.Fatal(e.Start(":8080"))
}