package html

import "io"

type IPAParams struct {
	Title   string
	Message string
}

func Ipa(w io.Writer, p IPAParams) error {
	return ipaHMTL.ExecuteTemplate(w, "layout.html", p)
}
