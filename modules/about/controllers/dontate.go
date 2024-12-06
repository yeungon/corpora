package about

import (
	"context"
	"log"
	"net/http"

	"github.com/yeungon/corpora/html/view"
	sqlite "github.com/yeungon/corpora/internal/database"
	aboutmodels "github.com/yeungon/corpora/modules/about/models"
)

func (app *Controller) Donate(w http.ResponseWriter, r *http.Request) {
	// Insert the data into the database
	db := sqlite.DB()
	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Printf("Database connection failed: %v", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	donationList := aboutmodels.GetAllDonation(db)
	p := view.DonateParams{
		Title:      "Vietnamese Corpora",
		Message:    "This is a new beginning! Hello from Index",
		DonateData: donationList,
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
