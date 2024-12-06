package aboutmodels

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	sqlite "github.com/yeungon/corpora/internal/database"
	"github.com/yeungon/corpora/internal/database/sqlite/donate"
)

func DonateUpdate(w http.ResponseWriter, r *http.Request) {
	var data donate.DonateData

	// Parse the incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Determine transaction type and amounts
	if data.TransferType == "in" {
		data.AmountIn = data.TransferAmount
	} else if data.TransferType == "out" {
		data.AmountOut = data.TransferAmount
	}
	// Insert the data into the database
	db := sqlite.DB()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		log.Printf("Database connection failed: %v", err)
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	_, err := db.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		log.Printf("Failed to insert record: %v", err)
		log.Printf("Database insertion failed: %v\nData: %+v", err, data)
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
