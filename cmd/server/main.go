package main

import (
	"log"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/app"
	"github.com/noworldwar/bc-service-livestream-api-go/internal/srvclient"
)

func main() {
	app.InitConfig()
	app.GenerateDoc()
	app.InitRedis()
	srvclient.Init()
	app.InitRouter()
	go app.RunRouter()

	// Add health check
	health := healthcheck.NewHandler()
	go func() {
		err := http.ListenAndServe("0.0.0.0:3000", health)
		if err != nil {
			log.Println("Health Error:", err)
		}
	}()

	app.Cleanup()
}
