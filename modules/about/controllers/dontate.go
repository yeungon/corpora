package about

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
)

func (app *Controller) Donate(w http.ResponseWriter, r *http.Request) {
	p := view.AboutParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	view.Donate(w, p)
}
