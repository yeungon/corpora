package aboutmodels

import (
	"context"
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/yeungon/corpora/internal/database/sqlite/donate"
)

func GetAllDonation(db *bun.DB) []donate.DonateData {
	var ctx = context.Background()
	var donateList []donate.DonateData

	// Debugging: Log query execution
	err := db.NewSelect().
		Model(&donateList).
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No data of donation found")
			return []donate.DonateData{}
		}
		log.Fatalf("Failed to retrieve donation: %v", err)
	}

	log.Printf("Retrieved donations: %+v", donateList)
	return donateList
}
