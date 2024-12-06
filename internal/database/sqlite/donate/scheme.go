package donate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
)

// DonateData represents the webhook data with Bun ORM.
type DonateData struct {
	bun.BaseModel      `bun:"table:donate"`
	ID                 int64     `bun:",pk,autoincrement"`
	Gateway            string    `bun:",notnull" json:"gateway"`
	TransactionDate    string    `bun:",notnull" json:"transactionDate"`
	AccountNumber      string    `bun:",notnull" json:"accountNumber"`
	SubAccount         *string   `bun:",nullzero" json:"subAccount"`
	TransferType       string    `bun:",notnull" json:"transferType"` // Match JSON key
	TransferAmount     float64   `bun:",notnull" json:"transferAmount"`
	AmountIn           float64   `bun:",notnull,default:0"`
	AmountOut          float64   `bun:",notnull,default:0"`
	Accumulated        float64   `bun:",nullzero,default:0" json:"accumulated"`
	Code               *string   `bun:",nullzero" json:"code"`
	TransactionContent string    `bun:",nullzero" json:"content"`
	ReferenceNumber    string    `bun:",nullzero" json:"referenceCode"`
	Body               string    `bun:",nullzero" json:"description"`
	CreatedAt          time.Time `bun:",notnull,default:current_timestamp"`
	UpdatedAt          time.Time `bun:",notnull,default:current_timestamp"`
}

func CreateTable(db *bun.DB) {
	var ctx = context.Background()
	res, err := db.NewCreateTable().Model((*DonateData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	info := fmt.Sprintf("The table 'donate' created. Rows affected: %d", res)
	fmt.Println(info)
}
