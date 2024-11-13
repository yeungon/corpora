package html

import (
	"io"
)

type IndexParams struct {
	Title       string
	Message     string
	StateSearch string
}

func Home(w io.Writer, p IndexParams) error {
	return home.ExecuteTemplate(w, "layout.html", p)
}
