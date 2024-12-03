package user

import (
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func (app *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	p := html.SignupUserParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	html.SignupUser(w, p, Partial(r))
}

func (app *Controller) ProfileShow(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	html.ProfileShow(w, p, Partial(r))
}

func (app *Controller) PhonemizerCtrl(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	html.PhonemizerHandle(w, p, Partial(r))
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
