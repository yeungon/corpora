package boot

import (
	"github.com/go-chi/chi/v5"
)

func Register(r *chi.Mux) {
	Router(r)
}
