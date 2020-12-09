package app

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/bc-service-livestream-api/internal/model"
	"github.com/spf13/viper"
)

func InitRouter() {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(Logger())
	r.Use(gin.Recovery())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad Request"})
	})

	log.Println("BC_API_PORT:" + viper.GetString("api_port"))

	model.APIServer = http.Server{
		Addr:    ":" + viper.GetString("api_port"),
		Handler: r,
	}
}

func RunRouter() {
	if err := model.APIServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
