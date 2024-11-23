package home

import (
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func (app *Controller) Credit(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	html.Credit(w, p)
}
