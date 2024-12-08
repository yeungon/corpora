package features

import (
	"fmt"
	"net/http"

	"github.com/yeungon/corpora/html/view"
	featuresmodels "github.com/yeungon/corpora/modules/features/models"
)

func (app *Controller) VietnameseIPA(w http.ResponseWriter, r *http.Request) {
	p := view.IPAParams{
		Title:   "International Phonetic Alphabet - ngữ âm tiếng Việt",
		Message: "Hello from IPA",
	}

	vietnamese := featuresmodels.VietnameseWord()
	fmt.Println(string(vietnamese))
	view.Ipa(w, p)
}
