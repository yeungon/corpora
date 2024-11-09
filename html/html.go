package html

import (
	"io"
)

var (
	index       = parse("template/index.html")
	profileShow = parse("template/profile/show.html")
	profileEdit = parse("template/profile/edit.html")
)

type IndexParams struct {
	Title   string
	Message string
}

func Index(w io.Writer, p IndexParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return index.ExecuteTemplate(w, partial, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}
