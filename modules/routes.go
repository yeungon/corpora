package modules

import (
	"github.com/go-chi/chi/v5"
	"github.com/yeungon/corpora/internal/config"
	about "github.com/yeungon/corpora/modules/about/controllers"
	home "github.com/yeungon/corpora/modules/home/controllers"
	user "github.com/yeungon/corpora/modules/user/controllers"
)

func RouteProvider(r *chi.Mux, appconfig *config.AppConfig) {
	homeCtrl := home.New(appconfig)
	r.Get("/", homeCtrl.Home)
	r.Get("/query", homeCtrl.SearchManticore)

	// User router
	userCtrl := user.New(appconfig)
	r.Get("/tokenize", userCtrl.ProfileShow)
	r.Get("/phonemizer", userCtrl.PhonemizerCtrl)
	r.Get("/signup", userCtrl.Signup)

	// About router
	aboutCtrl := about.New(appconfig)
	r.Get("/about", aboutCtrl.Introduction)
	r.Get("/credit", aboutCtrl.Credit)
}
