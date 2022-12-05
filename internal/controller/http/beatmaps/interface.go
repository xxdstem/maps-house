package beatmaps

import (
	"github.com/fasthttp/router"
)

type Handler interface {
	Register(router *router.Router)
}

type UseCase interface {
}
