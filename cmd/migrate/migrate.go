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

	err := db.AutoMigrate(&entity.BeatmapMeta{})
	fmt.Println(err)
	err = db.AutoMigrate(&entity.Beatmap{})
	fmt.Println(err)
	// Manually add foreign key constraint
	err = db.Exec("ALTER TABLE beatmaps ADD CONSTRAINT fk_beatmapset FOREIGN KEY (beatmapset_id) REFERENCES beatmap_meta(beatmapset_id)").Error
	fmt.Println(err)

	fmt.Println("Migration complete")
}
