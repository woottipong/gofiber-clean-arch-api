package main

import (
	"golang-clean-arch-api/config"
	"golang-clean-arch-api/internal/infrastructure"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initial API
	infrastructure.Initial(cfg)
}
