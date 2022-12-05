package app

import (
	"maps-house/config"
	"maps-house/internal/controller/http"
	"maps-house/internal/usecase"
	repo "maps-house/internal/usecase/repository/db"
	"maps-house/pkg/customrouter"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const listenAdress = ":8000"

func Run(conf *config.Config, log *logger.Logger) {

	// Build DSN string (probably could an another better way)
	dsn := conf.DSNBuilder()
	// Initiailize db
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Gorm error: %s", err)
	}
	// Initialize Router
	r := router.New()

	// Initialize custom router
	customRouter := customrouter.NewRouter(r)

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("index.")
	})
	// Initialize repos
	repo := repo.New(db)

	// Initialize usecases
	useCase := usecase.New(log, repo)

	// Initialize controllers
	http.NewApiRouter(r, log, useCase)
	log.Info("Listening app on ", listenAdress)
	err = fasthttp.ListenAndServe(listenAdress, customRouter.Handler)
	if err != nil {
		log.Fatal(err)
	}

}
