package home

import (
	"fmt"
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func (app *Controller) Home(w http.ResponseWriter, r *http.Request) {
	// test := config.GET().TEST
	// fmt.Println("Fetching config info: ", test)
	fmt.Println("In dữ liệu Dependency Injection: ", app.config.Test)

	p := html.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "This is a new beginning! Hello from Index",
	}
	//html.Home(w, p, Partial(r))
	html.Home(w, p)
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
