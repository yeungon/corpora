package modules

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
	home "github.com/yeungon/corpora/modules/home/controllers"
	user "github.com/yeungon/corpora/modules/user/controllers"
)

func RouteProvider(r *chi.Mux, appconfig *config.AppConfig) {
	homeCtrl := home.New(appconfig)
	r.Get("/", homeCtrl.Home)
	r.Post("/search/query", homeCtrl.SearchConcordancePost)

	// User router
	userCtrl := user.New(appconfig)
	r.Get("/profile", userCtrl.ProfileShow)
	r.Get("/profile/edit", userCtrl.ProfileEdit)
	r.Get("/signup", userCtrl.Signup)
}
