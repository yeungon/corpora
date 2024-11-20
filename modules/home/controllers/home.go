package home

import (
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func (app *Controller) Home(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	html.Home(w, p)
}
