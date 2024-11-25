package app

import (
	"maps-house/config"
	"maps-house/internal/controller/http"
	"maps-house/internal/services/beatmaps"
	"maps-house/internal/services/osuapi"
	"maps-house/internal/usecase"
	repo "maps-house/internal/usecase/repository/db"
	"maps-house/internal/workers/cleaner"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

const dataDogKey = "f983d1a4debdb5c7d46df7830d09fd5ddb125f82"

func Run(conf *config.Config, log *logger.Logger) {

	//statsd.New()
	// Build DSN string (probably could a better way)
	dsn := conf.DSNBuilder()

	// Initialize db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gLogger.Default.LogMode(gLogger.Silent),
	})

	if err != nil {
		log.Fatal("Gorm error: ", err)
	}
	// Initialize Router
	r := router.New()

	// Initialize custom router
	//customRouter := customrouter.NewRouter(r)

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		_, _ = ctx.WriteString("index.")
	})
	// Initialize repos
	dbRepo := repo.New(db)

	// Initialize services
	osuApiService := osuapi.NewService(log, conf.OsuApiKey)
	beatmapsService := beatmaps.NewService(log, conf.MapsPath)

	// Initialize useCases
	useCase := usecase.New(log, dbRepo, osuApiService, beatmapsService)

	_ = cleaner.New(log, db, beatmapsService)
	// Initialize controllers
	http.RegisterMain(r, log, useCase)
	http.RegisterApi(r, log, useCase)

	log.Info("Listening app on ", conf.ListenAddress)
	err = fasthttp.ListenAndServe(conf.ListenAddress, r.Handler)
	if err != nil {
		log.Fatal(err)
	}

}
