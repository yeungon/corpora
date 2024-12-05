package user

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
)

func (app *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	p := view.SignupUserParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	view.SignupUser(w, p, Partial(r))
}

func (app *Controller) ProfileShow(w http.ResponseWriter, r *http.Request) {
	p := view.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	view.ProfileShow(w, p, Partial(r))
}

func (app *Controller) PhonemizerCtrl(w http.ResponseWriter, r *http.Request) {
	p := view.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	view.PhonemizerHandle(w, p, Partial(r))
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
