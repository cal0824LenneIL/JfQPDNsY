// 代码生成时间: 2025-08-10 18:55:59
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
    "golang.org/x/crypto/pbkdf2"
)

// Configuration for AES encryption
type AESConfig struct {
    Key    string
    PBKDF2 IterationConfig
}

// IterationConfig for PBKDF2
type IterationConfig struct {
    Iterations int
    SaltLength int
}

// AESGCM is a struct to hold the AES configuration and key
type AESGCM struct {
    Key []byte
}

// NewAESGCM creates a new AESGCM instance with the given key
func NewAESGCM(key []byte) (*AESGCM, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    return &AESGCM{Key: key}, nil
}

// Encrypt encrypts the given plaintext using AES-GCM
func (a *AESGCM) Encrypt(plaintext []byte) (string, error) {
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    aesgcm, err := cipher.NewGCM(block.NewCipher(a.Key))
    if err != nil {
        return "", err
    }
    ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext using AES-GCM
func (a *AESGCM) Decrypt(ciphertext string) ([]byte, error) {
    decoded, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return nil, err
    }
    aesgcm, err := cipher.NewGCM(block.NewCipher(a.Key))
    if err != nil {
        return nil, err
    }
    nonceSize := aesgcm.NonceSize()
    if len(decoded) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
    return aesgcm.Open(nil, nonce, ciphertext, nil)
}

// generateKey generates a secure key for AES encryption
func generateKey(password string, config AESConfig) ([]byte, error) {
    salt := make([]byte, config.PBKDF2.SaltLength)
    if _, err := io.ReadFull(rand.Reader, salt); err != nil {
        return nil, err
    }
    key := pbkdf2.Key([]byte(password), salt, config.PBKDF2.Iterations, 32, sha256.New)
    return key, nil
}

func main() {
    e := echo.New()
    aesConfig := AESConfig{
        Key: "your-secret-key",
        PBKDF2: IterationConfig{
            Iterations: 10000,
            SaltLength: 16,
        },
    }
    key, err := generateKey(aesConfig.Key, aesConfig)
    if err != nil {
        log.Fatal(err)
    }
    aesGCM, err := NewAESGCM(key)
    if err != nil {
        log.Fatal(err)
    }
    
    // Encrypt endpoint
    e.POST="/encrypt", func(c echo.Context) error {
        var data struct {
            Plaintext string `json:"plaintext"`
        }
        if err := c.Bind(&data); err != nil {
            return err
        }
        encrypted, err := aesGCM.Encrypt([]byte(data.Plaintext))
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{"encrypted": encrypted})
    })
    
    // Decrypt endpoint
    e.POST="/decrypt", func(c echo.Context) error {
        var data struct {
            Ciphertext string `json:"ciphertext"`
        }
        if err := c.Bind(&data); err != nil {
            return err
        }
        decrypted, err := aesGCM.Decrypt(data.Ciphertext)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, map[string]string{"decrypted": string(decrypted)})
    })
    
    e.Logger.Fatal(e.Start(":8080"))
}