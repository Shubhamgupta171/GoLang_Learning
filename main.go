package main

import (
	"go-api/controllers"
	"go-api/middleware"
	"go-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database + Redis
	database.Connect()
	database.ConnectRedis()

	r := gin.Default()

	// Public
	r.POST("/auth/register", controllers.Register)

	r.POST("/auth/login",
		middleware.RedisRateLimiter(10, 10.0/60000.0, "login"),
		controllers.Login)

	r.POST("/auth/refresh",
		middleware.RedisRateLimiter(10, 10.0/60000.0, "refresh"),
		controllers.RefreshToken)

	// Logout (must have route)
	r.POST("/auth/logout",
		middleware.RedisRateLimiter(10, 10.0/60000.0, "logout"),
		controllers.Logout)

	// Protected
	protected := r.Group("/")
	protected.Use(
		middleware.AuthMiddleware(),
		middleware.RedisRateLimiter(5, 5.0/60000.0, "profile"),
	)

	{
		protected.GET("/profile", controllers.GetProfile)
	}

	r.Run(":8080")
}
