package boot

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func RegisterStatic(r *chi.Mux) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fmt.Println(filesDir)
	fileServer := http.StripPrefix("/static", http.FileServer(http.Dir(filesDir)))
	r.Handle("/static/*", fileServer)
}

//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
