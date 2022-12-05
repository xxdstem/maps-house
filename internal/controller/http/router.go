package http

import (
	"maps-house/internal/controller/http/api"
	"maps-house/internal/usecase"
	"maps-house/pkg/logger"

	"github.com/fasthttp/router"
)

func NewApiRouter(r *router.Router, l *logger.Logger, uc usecase.UseCase) {
	apiHandler := api.New(uc, l)
	apiHandler.Register(r)
}
