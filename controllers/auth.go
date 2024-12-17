package controllers

import (
    "errors"
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/yourusername/ai-backend/models"
)

var jwtSecret = []byte("your_secret_key")

// GenerateToken 生成 JWT
func GenerateToken(user models.User) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &jwt.StandardClaims{
        Subject:   fmt.Sprintf("%d", user.ID),
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*jwt.Token, error) {
    return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
}

// AuthMiddleware 中间件：检查请求中的 JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(400, gin.H{"error": "Authorization token not provided"})
            c.Abort()
            return
        }

        _, err := ParseToken(tokenString)
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}