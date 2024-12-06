package donate

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

func GetAllUser(db *bun.DB) []WebhookData {
	var ctx = context.Background()
	// Retrieve all users.
	var donateList []WebhookData
	err := db.NewSelect().
		Model(&donateList).
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found")
			return []WebhookData{} // Return a zero-value User if no user is found.
		}
		log.Fatal("Failed to retrieve users:", err)
	}

	return donateList
}
