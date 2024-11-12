package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
)

func Register(r *chi.Mux, appconfig *config.AppConfig) {
	// Router(r)
	RegisterRouter(r, appconfig)
	Static(r)
	config.New()
}
