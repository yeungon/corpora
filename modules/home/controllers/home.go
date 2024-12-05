package home

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
)

func (app *Controller) Home(w http.ResponseWriter, r *http.Request) {
	p := view.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	view.Home(w, p)
}
