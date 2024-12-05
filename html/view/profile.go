package view

import (
	"io"

	"github.com/yeungon/corpora/html"
)

type SignupUserParams struct {
	Title   string
	Message string
}

func SignupUser(w io.Writer, p SignupUserParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return html.PageSignupUser.ExecuteTemplate(w, partial, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func ProfileShow(w io.Writer, p ProfileShowParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return html.PageTokenize.ExecuteTemplate(w, partial, p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func PhonemizerHandle(w io.Writer, p ProfileEditParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return html.Pagephonemizer.ExecuteTemplate(w, partial, p)
}
