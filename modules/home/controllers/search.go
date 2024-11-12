package home

import (
	"net/http"
)

func (app *Controller) SearchConcordancePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	w.Write([]byte("welcome"))

}
