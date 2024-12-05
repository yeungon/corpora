package about

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
)

func (app *Controller) Credit(w http.ResponseWriter, r *http.Request) {
	p := view.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	view.Credit(w, p)
}
