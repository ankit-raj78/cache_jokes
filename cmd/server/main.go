package main

import (
	"log"

	"github.com/yourorg/cache_jokes/internal/api"
	"github.com/yourorg/cache_jokes/internal/config"
)

func main() {
	// 1. Load config (e.g. port)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2. Create router & start HTTP server
	router := api.NewRouter()
	log.Printf("starting server on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
