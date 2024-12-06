package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
	sqlite "github.com/yeungon/corpora/internal/database"
)

func Register(r *chi.Mux, appconfig *config.AppConfig) {
	// Router(r)
	RegisterRouter(r, appconfig)
	Static(r)
	config.New()
	// =============sqlite=============
	sqlite.BunConnect()
	db := sqlite.DB()
	defer db.Close()
	// Enable debugging for queries
	// =============sqlite=============
}
