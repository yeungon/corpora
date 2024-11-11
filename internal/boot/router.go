package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/modules/home"
)

func Router(r *chi.Mux) {
	r.Get("/", home.Home)
	r.Get("/profile", ProfileShow)
	r.Get("/profile/edit", ProfileEdit)
	r.Get("/profile/show", ProfileShow)
}
