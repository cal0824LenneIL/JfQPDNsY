// 代码生成时间: 2025-08-28 05:33:28
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "github.com/labstack/echo"
)

// CryptoTool 结构体用于加密和解密
type CryptoTool struct {
    key []byte
}

// NewCryptoTool 初始化CryptoTool实例
func NewCryptoTool(key []byte) *CryptoTool {
    return &CryptoTool{key: key}
}

// Encrypt 加密函数
func (ct *CryptoTool) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(ct.key)
    if err != nil {
        return "", err
    }

    // 填充
    plaintextBytes := PKCS7Padding([]byte(plaintext), aes.BlockSize)
    // 加密
    blockMode := cipher.NewCBCEncrypter(block, ct.key[:block.BlockSize()])
    encryptedBytes := make([]byte, len(plaintextBytes))
    blockMode.CryptBlocks(encryptedBytes, plaintextBytes)

    // 编码
    return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// Decrypt 解密函数
func (ct *CryptoTool) Decrypt(ciphertext string) (string, error) {
    encryptedBytes, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(ct.key)
    if err != nil {
        return "", err
    }

    blockMode := cipher.NewCBCDecrypter(block, ct.key[:block.BlockSize()])
    decryptedBytes := make([]byte, len(encryptedBytes))
    blockMode.CryptBlocks(decryptedBytes, encryptedBytes)

    // 去除填充
    decryptedBytes = PKCS7UnPadding(decryptedBytes)
    return string(decryptedBytes), nil
}

// PKCS7Padding 填充函数
func PKCS7Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padText...)
}

// PKCS7UnPadding 去除填充函数
func PKCS7UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length-1])
    return src[:(length - unpadding)]
}

func main() {
    key := []byte("your-32-byte-key") // 密钥长度必须是32字节
    cryptoTool := NewCryptoTool(key)

    e := echo.New()
    
    // 加密接口
    e.POST("/encrypt", func(c echo.Context) error {
        plaintext := c.FormValue("plaintext")
        encrypted, err := cryptoTool.Encrypt(plaintext)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{
            "encrypted": encrypted,
        })
    })
    
    // 解密接口
    e.POST("/decrypt", func(c echo.Context) error {
        ciphertext := c.FormValue("ciphertext")
        decrypted, err := cryptoTool.Decrypt(ciphertext)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{
            "decrypted": decrypted,
        })
    })

    // 启动服务
    e.Logger.Fatal(e.Start(":1323"))
}
