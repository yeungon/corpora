package features

import (
	"fmt"
	"net/http"

	"github.com/yeungon/corpora/html/view"
	featuresmodels "github.com/yeungon/corpora/modules/features/models"
)

func (app *Controller) EnglishIPA(w http.ResponseWriter, r *http.Request) {
	p := view.IPAParams{
		Title:   "International Phonetic Alphabet - ngữ âm tiếng Việt",
		Message: "Hello from IPA",
	}

	// Example paragraph
	paragraph := "'cause 'tis the season to be merry, and 'bout time we did something fun!"

	// Get pronunciations
	// Get pronunciations
	pronunciations, err := featuresmodels.EnglishWord(paragraph)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print results
	fmt.Println("Word Pronunciations:")
	for word, pronunciation := range pronunciations {
		fmt.Printf("%s: %s\n", word, pronunciation)
	}
	view.Ipa(w, p)
}
