package main

import (
	"maps-house/config"
	"maps-house/internal/app"
	"maps-house/pkg/logger"
	"os"
)

func main() {
	// Configuration
	logger.Init()
	log := logger.New()
	conf, err := config.NewConfig(log)

	if err != nil {
		log.Fatal("Config error: ", err)
		os.Exit(-1)
	}

	// Run
	app.Run(conf, log)
}
