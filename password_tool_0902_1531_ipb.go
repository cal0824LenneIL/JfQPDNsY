// 代码生成时间: 2025-09-02 15:31:54
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/pbkdf2"
    "net/http"
    "time"
    "github.com/labstack/echo"
)

const (
    key = "your-secure-key" // 密钥，应替换为安全的环境变量值
    nonce = "your-secure-nonce" // 随机数（nonce），应替换为安全的环境变量值
)

// 加密函数
func encrypt(plaintext string) (string, error) {
    k := []byte(key)
    n := []byte(nonce)
    ciphertext := pbkdf2.Key(k, n, 4096, 32, sha256.New)
    block, err := aes.NewCipher(ciphertext)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := []byte(n)
    encrypted := gcm.Seal(nonce, nonce[:12], []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 解密函数
func decrypt(ciphertext string) (string, error) {
    k := []byte(key)
    n := []byte(nonce)
    ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    ciphertextNonce := ciphertextBytes[:12]
    ciphertextBytes = ciphertextBytes[12:]
    
    ciphertext = pbkdf2.Key(k, n, 4096, 32, sha256.New)
    block, err := aes.NewCipher(ciphertext)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    if len(ciphertextNonce) < gcm.NonceSize() {
        return "", fmt.Errorf("nonce is too short")
    }
    nonce := ciphertextNonce[:gcm.NonceSize()]
    return string(gcm.Open(nil, nonce, ciphertextBytes, nil)), nil
}

// 处理加密请求
func encryptHandler(c echo.Context) error {
    input := c.QueryParam("text")
    encrypted, err := encrypt(input)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(http.StatusOK, echo.Map{
        "original": input,
        "encrypted": encrypted,
    })
}

// 处理解密请求
func decryptHandler(c echo.Context) error {
    encrypted := c.QueryParam("text")
    decrypted, err := decrypt(encrypted)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(http.StatusOK, echo.Map{
        "encrypted": encrypted,
        "decrypted": decrypted,
    })
}

func main() {
    e := echo.New()
    
    // 加密路由
    e.GET("/encrypt", encryptHandler)
    // 解密路由
    e.GET("/decrypt", decryptHandler)
    
    e.Logger.Fatal(e.Start(":8080"))
}