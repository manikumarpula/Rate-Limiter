package main

import (
	"log"
	"rate-limiter/api/http"
	"rate-limiter/configs"
)

func main() {
	log.Println("Rate Limiter service starting...")
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("could not load configuration: %v", err)
	}
	http.StartServer(&cfg)
}
