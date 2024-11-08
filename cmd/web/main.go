package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	web "github.com/yeungon/corpora/cmd"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to Vietnamese Corpora"))
	})
	r.Get("/index", web.Index)
	r.Get("/profile", web.ProfileShow)
	r.Get("/profile/edit", web.ProfileEdit)
	http.ListenAndServe(":9999", r)
}
