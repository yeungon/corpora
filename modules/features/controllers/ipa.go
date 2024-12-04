package features

import (
	"net/http"

	"github.com/yeungon/corpora/html"
)

func (app *Controller) IPA(w http.ResponseWriter, r *http.Request) {
	p := html.IPAParams{
		Title:   "International Phonetics Alphabet - ngữ âm tiếng Việt",
		Message: "Hello from profile show",
	}

	html.Ipa(w, p)
}
