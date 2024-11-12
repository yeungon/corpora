package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
	"github.com/yeungon/corpora/modules"
)

func RegisterRouter(r *chi.Mux, appconfig *config.AppConfig) {
	modules.RouteProvider(r, appconfig)
}
