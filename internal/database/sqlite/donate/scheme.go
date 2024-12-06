package donate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

// WebhookData represents the webhook data with Bun ORM.
type WebhookData struct {
	bun.BaseModel      `bun:"table:donate"` // Specifies the table name for the Bun ORM.
	ID                 int64                `bun:",pk,autoincrement"`                  // Primary key with auto-increment.
	Gateway            string               `bun:",notnull"`                           // Payment gateway identifier.
	TransactionDate    time.Time            `bun:",notnull"`                           // Date of the transaction.
	AccountNumber      string               `bun:",notnull"`                           // Account number involved in the transaction.
	SubAccount         string               `bun:",nullzero"`                          // Sub-account information, optional.
	AmountIn           float64              `bun:",notnull,default:0"`                 // Amount received in the transaction.
	AmountOut          float64              `bun:",notnull,default:0"`                 // Amount sent out in the transaction.
	Accumulated        float64              `bun:",nullzero,default:0"`                // Accumulated balance after the transaction.
	Code               string               `bun:",notnull"`                           // Transaction code.
	TransactionContent string               `bun:",nullzero"`                          // Transaction description or content.
	ReferenceNumber    string               `bun:",nullzero"`                          // Reference number for the transaction.
	Body               string               `bun:",nullzero"`                          // Additional details about the transaction.
	CreatedAt          time.Time            `bun:",notnull,default:current_timestamp"` // Record creation timestamp.
	UpdatedAt          time.Time            `bun:",notnull,default:current_timestamp"` // Record update timestamp.
}

func CreateTable(db *bun.DB) {
	var ctx = context.Background()
	res, err := db.NewCreateTable().Model((*WebhookData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	info := fmt.Sprintf("The table 'users' created. Rows affected: %d", res)
	fmt.Println(info)
}
