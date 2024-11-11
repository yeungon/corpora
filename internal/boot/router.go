package boot

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/modules/home"
	"github.com/yeungon/corpora/modules/user"
)

func Router(r *chi.Mux) {
	r.Get("/", home.Home)
	r.Get("/profile", user.ProfileShow)
	r.Get("/profile/edit", user.ProfileEdit)
}
