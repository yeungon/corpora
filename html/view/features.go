package view

import (
	"io"

	"github.com/yeungon/corpora/html"
)

type IPAParams struct {
	Title   string
	Message string
}

func Ipa(w io.Writer, p IPAParams) error {
	return html.PageIPA.ExecuteTemplate(w, "layout.html", p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func PhonemizerView(w io.Writer, p ProfileEditParams) error {
	return html.Pagephonemizer.ExecuteTemplate(w, "layout.html", p)
}
