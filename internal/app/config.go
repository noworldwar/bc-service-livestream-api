package app

import (
	"fmt"
	"os"

	"github.com/micro/go-micro/config/cmd"
	"github.com/spf13/viper"
)

func InitConfig() {
	err := cmd.Init()
	if err != nil {
		fmt.Println("go-micro cmd Init Error:", err)
	}

	// set prefix for env
	viper.SetEnvPrefix("bc")
	// set default
	viper.SetDefault("api_port", "8080")
	viper.SetDefault("redis_hostname", "localhost")
	viper.SetDefault("redis_port", "6379")
	viper.SetDefault("redis_cluster", "")
	viper.SetDefault("redis_database", 1)
	viper.SetDefault("redis_expiration", 1) // hour

	viper.AutomaticEnv()

	fmt.Printf("[Config] MICRO_REGISTRY: %v \r\n", os.Getenv("MICRO_REGISTRY"))
	fmt.Printf("[Config] MICRO_REGISTRY_ADDRESS: %v \r\n", os.Getenv("MICRO_REGISTRY_ADDRESS"))
}
