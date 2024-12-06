package about

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
	aboutmodels "github.com/yeungon/corpora/modules/about/models"
)

func (app *Controller) Donate(w http.ResponseWriter, r *http.Request) {
	p := view.AboutParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	view.Donate(w, p)
}

func (app *Controller) ReceiveDonate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Page not found!"))
		return
	}

	// Restrict to POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	aboutmodels.DonateUpdate(w, r)
}
