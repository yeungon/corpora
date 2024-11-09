package boot

import (
	"github.com/go-chi/chi/v5"
)

func Router(r *chi.Mux) {
	r.Get("/", Index)
	r.Get("/profile", ProfileShow)
	r.Get("/profile/edit", ProfileEdit)
	r.Get("/profile/show", ProfileShow)
}
