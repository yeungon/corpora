package view

import (
	"io"

	"github.com/yeungon/corpora/html"
)

type DonateParams struct {
	Title   string
	Message string
}

func Donate(w io.Writer, p AboutParams) error {
	return html.PageDonate.ExecuteTemplate(w, "layout.html", p)
}
