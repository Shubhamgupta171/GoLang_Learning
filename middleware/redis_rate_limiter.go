package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/database"
)

// Atomic Lua Script - Token Bucket
const tokenBucketLua = `
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])

local data = redis.call('HMGET', key, 'tokens', 'last')
local tokens = tonumber(data[1])
local last = tonumber(data[2])

if tokens == nil then
	tokens = capacity
	last = now
end

local delta = now - last
local add = delta * rate
tokens = math.min(capacity, tokens + add)
last = now

local allowed = tokens >= requested
if allowed then
	tokens = tokens - requested
	redis.call('HMSET', key, 'tokens', tokens, 'last', last)
	redis.call('EXPIRE', key, math.ceil(capacity / rate))
	return {1, tokens}
else
	redis.call('HMSET', key, 'tokens', tokens, 'last', last)
	redis.call('EXPIRE', key, math.ceil(capacity / rate))
	return {0, tokens}
end
`

// capacity: number of requests allowed in bucket
// rate: tokens added per millisecond (0.1 = 1 token every 10ms)
func RedisRateLimiter(capacity int, rate float64, keyPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.ClientIP()
		key := fmt.Sprintf("rate:%s:%s", keyPrefix, clientIP)

		now := time.Now().UnixMilli()

		res, err := database.RDB.Eval(database.RdbCtx, tokenBucketLua,
			[]string{key}, capacity, rate, now, 1).Result()

		if err != nil {
			fmt.Println("Redis Error:", err)
			c.Next()
			return
		}

		data := res.([]interface{})
		allowed := data[0].(int64)

		if allowed == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
