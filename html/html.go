package html

import (
	"embed"
	"io"
)

//go:embed *
var files embed.FS

var (
	index       = parse("index.html")
	profileShow = parse("profile/show.html")
	profileEdit = parse("profile/edit.html")
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

func ProfileShow(w io.Writer, p ProfileShowParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return profileShow.ExecuteTemplate(w, partial, p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func ProfileEdit(w io.Writer, p ProfileEditParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return profileEdit.ExecuteTemplate(w, partial, p)
}
