package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
)

func Register(r *chi.Mux) {
	Router(r)
	Static(r)
	config.New()
}
