package html

import (
	"io"
)

type IndexParams struct {
	Title   string
	Message string
}

func Home(w io.Writer, p IndexParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return home.ExecuteTemplate(w, partial, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}
