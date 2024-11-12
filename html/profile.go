package html

import "io"

type SignupUserParams struct {
	Title   string
	Message string
}

func SignupUser(w io.Writer, p SignupUserParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return signupUser.ExecuteTemplate(w, partial, p)
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
