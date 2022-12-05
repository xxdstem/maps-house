package http

import (
	"mapsHouse/internal/controller/http/api"
	"mapsHouse/internal/usecase"
	"mapsHouse/pkg/logger"

	"github.com/fasthttp/router"
)

func NewApiRouter(r *router.Router, l *logger.Logger, uc usecase.UseCase) {
	apiHandler := api.New(l, uc)
	apiHandler.Register(r)
}
