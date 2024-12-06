package view

import (
	"io"

	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/internal/database/sqlite/donate"
)

type DonateParams struct {
	Title      string
	Message    string
	DonateData []donate.DonateData
}

func Donate(w io.Writer, p DonateParams) error {
	return html.PageDonate.ExecuteTemplate(w, "layout.html", p)
}
