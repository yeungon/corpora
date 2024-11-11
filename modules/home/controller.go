package home

import (
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func Home(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	//html.Home(w, p, Partial(r))
	html.Home(w, p)
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
