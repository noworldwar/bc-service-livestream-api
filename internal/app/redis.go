package app

import (
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/model"
	"github.com/spf13/viper"
)

func InitRedis() {
	// get cluster settings
	redisClusterAddr := viper.GetString("redis_cluster")

	// create redis client
	if redisClusterAddr == "" {
		redisURL := viper.GetString("redis_hostname") + ":" + viper.GetString("redis_port")
		db := viper.GetInt("redis_database")

		log.Printf("Redis Standalone Mode on: %v database: %v", redisURL, db)

		model.RedisDB = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       db, // use default DB
		})
	} else {
		// get the address
		clusterAddress := strings.Split(redisClusterAddr, ",")

		log.Printf("Redis Cluster Mode on: %v", redisClusterAddr)

		model.RedisDB = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: clusterAddress,
		})
	}

	model.RedisEX = time.Duration(viper.GetInt("redis_expiration")) * time.Hour
}
