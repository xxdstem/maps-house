package http

import (
	"maps-house/internal/controller/http/api"
	"maps-house/internal/controller/http/beatmaps"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
)

func RegisterApi(r *router.Router, l *logger.Logger, uc UseCase) {
	apiHandler := api.New(uc, l)
	apiHandler.Register(r)
}

func RegisterMain(r *router.Router, l *logger.Logger, uc UseCase) {
	mainHandler := beatmaps.New(uc, l)
	mainHandler.Register(r)
}
