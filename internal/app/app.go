package app

import (
	"mapsHouse/config"
	"mapsHouse/internal/controller/http"
	repo "mapsHouse/internal/repository/db"
	"mapsHouse/internal/usecase"
	"mapsHouse/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const listenAdress = ":8000"

func Run(conf *config.Config, log *logger.Logger) {
	// Build DSN string (probably could an another better way)
	dsn := conf.DSNBuilder()
	log.Info("test")
	// Initiailize db
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Gorm error: %s", err)
	}
	// Initialize Router
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		log.Info("test")
		ctx.WriteString("index.")
	})
	// Initialize repos
	repo := repo.New(db)

	// Initialize usecases
	useCase := usecase.New(log, repo)

	// Initialize controllers
	http.NewApiRouter(r, log, useCase)
	log.Info("Listening app on ", listenAdress)
	err = fasthttp.ListenAndServe(listenAdress, r.Handler)
	if err != nil {
		log.Fatal(err)
	}

}
