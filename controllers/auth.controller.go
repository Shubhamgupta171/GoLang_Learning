package controllers

import (
    "go-api/database"
    "go-api/models"
    "go-api/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

// ===================== REGISTER ========================
func Register(c *gin.Context) {
    var body models.RegisterRequest

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Hash password
    hashed, err := utils.HashPassword(body.Password)
    if err != nil {
        c.JSON(500, gin.H{"error": "Password hashing failed"})
        return
    }

    // Insert into DB
    _, err = database.DB.Exec(
        "INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
        body.Name, body.Email, hashed,
    )
    if err != nil {
        c.JSON(400, gin.H{"error": "Email already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// ===================== LOGIN =============================
func Login(c *gin.Context) {
    var body models.LoginRequest

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    // Fetch user
    var user models.User
    err := database.DB.Get(&user, "SELECT * FROM users WHERE email = ?", body.Email)
    if err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    // Validate password
    if !utils.CheckPasswordHash(body.Password, user.Password) {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate tokens
    accessToken, errA := utils.GenerateAccessToken(user.Email)
    refreshToken, errR := utils.GenerateRefreshToken(user.Email)

    if errA != nil || errR != nil {
        c.JSON(500, gin.H{"error": "Token generation failed"})
        return
    }

    // Return login success
    c.JSON(200, gin.H{
        "message":       "Login successful",
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
}
