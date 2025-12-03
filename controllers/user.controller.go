package controllers

import (
    "go-api/database"
    "go-api/models"
    "go-api/utils"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type ProfileResponse struct {
    Message string `json:"message"`
    Email   string `json:"email"`
    Name    string `json:"name"`
}

func GetProfile(c *gin.Context) {
    emailValue, exists := c.Get("email")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email missing in token"})
        return
    }

    email := emailValue.(string)
    cacheKey := "profile:" + email

   
    var cachedUser models.User
    found, _ := utils.CacheGet(cacheKey, &cachedUser)

    if found {
        c.JSON(200, ProfileResponse{
            Message: "Profile fetched",
            Email:   cachedUser.Email,
            Name:    cachedUser.Name,
        })
        return
    }

    
    var user models.User
    err := database.DB.Get(&user, "SELECT name, email FROM users WHERE email = ?", email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

  
    utils.CacheSet(cacheKey, user, 30*time.Second)


    c.JSON(200, ProfileResponse{
        Message: "Profile fetched",
        Email:   user.Email,
        Name:    user.Name,
    })
}
