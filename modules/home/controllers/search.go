package home

import (
	"fmt"
	"net/http"

	html "github.com/yeungon/corpora/html"
)

func (app *Controller) SearchConcordancePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	var form postCreateForm

	err := app.helper.DecodePostForm(r, &form)
	if err != nil {
		app.helper.ClientError(w, http.StatusBadRequest)
		return
	}

	fmt.Println("===================================in form")
	fmt.Println(form)
	fmt.Println("===================================")
	//w.Write([]byte("welcome"))
	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     "Đã search",
		StateSearch: "searched",
	}
	//html.Home(w, p, Partial(r))
	html.Home(w, p)

}
