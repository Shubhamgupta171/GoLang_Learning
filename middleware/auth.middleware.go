package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "go-api/database"
    "go-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        header := c.GetHeader("Authorization")

        if header == "" || !strings.HasPrefix(header, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(header, "Bearer ")

        // Check blacklist
        if database.BlacklistedTokens[tokenString] {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been logged out"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return utils.JwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Set("email", claims["email"])

        c.Next()
    }
}
