package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 2 requests per second, burst upto 5
var throttleLimiter = rate.NewLimiter(2, 5)

func Throttle() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Allow instantly if tokens available
		if throttleLimiter.Allow() {
			c.Next()
			return
		}

		// Otherwise reject extra requests
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "⚠️ Too many requests (throttled). Slow down!",
		})
		c.Abort()
	}
}
