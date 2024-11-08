package web

import (
	"net/http"

	"github.com/yeungon/corpora/html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := html.DashboardParams{
		Title:   "Vietnamese Corpora",
		Message: "Hello from Index \n",
	}
	html.Index(w, p, Partial(r))
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	p := html.DashboardParams{
		Title:   "Dashboard",
		Message: "Hello from dashboard",
	}
	html.Dashboard(w, p, Partial(r))
}

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
