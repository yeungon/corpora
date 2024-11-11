package boot

import (
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func ProfileShow(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	html.ProfileShow(w, p, Partial(r))
}

func ProfileEdit(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	html.ProfileEdit(w, p, Partial(r))
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
