package boot

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadStatic(r *chi.Mux) {
	fileServer := http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))
	r.Handle("/static/*", fileServer)
}

//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
