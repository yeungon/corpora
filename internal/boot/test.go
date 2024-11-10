package boot

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterStatic(r chi.Router) {
	fileServer := http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/")))
	r.Handle("/assets/*", fileServer)
}

//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
