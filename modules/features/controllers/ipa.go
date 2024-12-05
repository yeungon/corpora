package features

import (
	"net/http"

	"github.com/yeungon/corpora/html/view"
)

func (app *Controller) IPA(w http.ResponseWriter, r *http.Request) {
	p := view.IPAParams{
		Title:   "International Phonetics Alphabet - ngữ âm tiếng Việt",
		Message: "Hello from IPA",
	}

	view.Ipa(w, p)
}
