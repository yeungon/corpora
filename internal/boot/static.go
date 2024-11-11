package boot

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func LoadStatic(r *chi.Mux) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)
}

func FileServer(r *chi.Mux, path string, root http.FileSystem) {
	// Ensure there are no URL parameters
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	// Serve static files at the given path
	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(path)
		contentType := mime.TypeByExtension(ext)
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		fs := http.FileServer(root)
		fs.ServeHTTP(w, r)
	})

}

//https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
