package about

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/yeungon/corpora/html/view"
	sqlite "github.com/yeungon/corpora/internal/database"
	"github.com/yeungon/corpora/internal/database/sqlite/donate"
)

func (app *Controller) Donate(w http.ResponseWriter, r *http.Request) {
	p := view.AboutParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	view.Donate(w, p)
}

func (app *Controller) ReceiveDonate(w http.ResponseWriter, r *http.Request) {

	var data donate.WebhookData

	// Parse the incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Determine transaction type and amounts
	if data.Code == "in" {
		data.AmountIn = data.Accumulated
	} else if data.Code == "out" {
		data.AmountOut = data.Accumulated
	}

	// Insert the data into the database
	db := sqlite.DB()

	ctx := context.Background()

	_, err := db.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		http.Error(w, "Database insertion failed", http.StatusInternalServerError)
		log.Printf("Failed to insert record: %v", err)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})

	// p := view.AboutParams{
	// 	Title:   "Vietnamese Corpora",
	// 	Message: "Update",
	// }
	// view.Donate(w, p)
}
