package boot

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func LoadStatic(r *chi.Mux) {
	log.Println("Loading static files handler...")
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fmt.Println("Static files are served from:", filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)
}

func FileServer(r *chi.Mux, path string, root http.FileSystem) {
	// Ensure there are no URL parameters
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	// Add trailing slash if missing, ensuring correct handling of the directory
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	// Serve static files at the given path
	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		// Log the requested URL path to help with debugging
		log.Printf("Request URL: %s", r.URL.Path)

		ext := filepath.Ext(path)
		contentType := mime.TypeByExtension(ext)

		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}

		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")

		// Use http.StripPrefix to serve files correctly
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))

		log.Printf("Serving file: %s", r.URL.Path)
		fs.ServeHTTP(w, r)
	})

}

//https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//Another (sound) comprehensive solution, for future reference: https://github.com/go-chi/chi/issues/155
