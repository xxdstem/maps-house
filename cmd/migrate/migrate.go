package main

import (
	"fmt"
	"maps-house/config"
	"maps-house/internal/entity"
	"maps-house/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

var db *gorm.DB
var log *logger.Logger

func init() {
	// Configuration
	logger.Init()
	log = logger.New()
	conf, err := config.NewConfig(log)
	if err != nil {
		log.Fatal("Config error: ", err)
	}
	dsn := conf.DSNBuilder()

	// Initialize db
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gLogger.Default.LogMode(gLogger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db.AutoMigrate(&entity.Beatmap{})
	db.AutoMigrate(&entity.BeatmapMeta{})
	log.Done("? Migration complete")
	fmt.Println("???")
}
