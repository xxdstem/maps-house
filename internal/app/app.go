package app

import (
	"maps-house/config"
	"maps-house/internal/controller/http"
	"maps-house/internal/services/beatmaps"
	"maps-house/internal/services/osuapi"
	"maps-house/internal/usecase"
	repo "maps-house/internal/usecase/repository/db"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

func Run(conf *config.Config, log *logger.Logger) {

	// Build DSN string (probably could a better way)
	dsn := conf.DSNBuilder()

	// Initialize db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gLogger.Default.LogMode(gLogger.Silent),
	})

	if err != nil {
		log.Fatal("Gorm error: %s", err)
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
	beatmapsService := beatmaps.NewService(log, dbRepo, conf.Dirs.PriorityDir, conf.Dirs.MainDir)

	// Initialize useCases
	useCase := usecase.New(log, dbRepo, osuApiService, beatmapsService)

	// Initialize controllers
	http.RegisterMain(r, log, useCase)
	http.RegisterApi(r, log, useCase)

	log.Info("Listening app on ", conf.ListenAddress)
	err = fasthttp.ListenAndServe(conf.ListenAddress, r.Handler)
	if err != nil {
		log.Fatal(err)
	}

}
