package boot

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func Static(r *chi.Mux) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer := http.StripPrefix("/static", http.FileServer(http.Dir(filesDir)))
	r.Handle("/static/*", fileServer)
}

//https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
