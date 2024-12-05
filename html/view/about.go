package view

import (
	"io"

	"github.com/yeungon/corpora/html"
)

type AboutParams struct {
	Title   string
	Message string
}

func About(w io.Writer, p AboutParams) error {
	return html.PageAbout.ExecuteTemplate(w, "layout.html", p)
}
