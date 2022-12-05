package main

import (
	"mapsHouse/config"
	"mapsHouse/internal/app"
	"mapsHouse/pkg/logger"
)

func main() {
	// Configuration
	logger.Init()
	log := logger.New()
	conf, err := config.NewConfig(log)

	if err != nil {
		log.Fatal("Config error: %s", err)
	}

	// Run
	app.Run(conf, log)
}
