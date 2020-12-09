package model

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

// RedisClient redis client interface
type RedisClient interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Del(...string) *redis.IntCmd
	Keys(pattern string) *redis.StringSliceCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
}

var APIServer http.Server
var RedisDB RedisClient
var RedisEX time.Duration
