package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	web "github.com/yeungon/corpora/cmd"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/dashboard", web.Index)
	r.Get("/profile/show", web.ProfileShow)
	r.Get("/profile/edit", web.ProfileEdit)
	http.ListenAndServe(":9999", r)
}
