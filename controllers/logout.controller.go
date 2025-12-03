package controllers

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "go-api/database"
)

func Logout(c *gin.Context) {
    header := c.GetHeader("Authorization")

    if header == "" || !strings.HasPrefix(header, "Bearer ") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
        return
    }

    token := strings.TrimPrefix(header, "Bearer ")

    // Add to blacklist
    database.BlacklistedTokens[token] = true

    c.JSON(http.StatusOK, gin.H{
        "message": "Logout successful",
    })
}
