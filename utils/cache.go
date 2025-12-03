package utils

import (
	"encoding/json"
	"time"

	"go-api/database"
)

func CacheSet(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return database.RDB.Set(database.RdbCtx, key, data, ttl).Err()
}

func CacheGet(key string, dest interface{}) (bool, error) {
	data, err := database.RDB.Get(database.RdbCtx, key).Result()
	if err != nil {
		return false, nil // cache miss
	}
	return true, json.Unmarshal([]byte(data), dest)
}

func CacheDelete(key string) {
	database.RDB.Del(database.RdbCtx, key)
}
