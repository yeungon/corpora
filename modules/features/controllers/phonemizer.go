package features

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yeungon/corpora/html/view"
	featuresmodels "github.com/yeungon/corpora/modules/features/models"
)

func (app *Controller) PhonemizerCtrl(w http.ResponseWriter, r *http.Request) {
	p := view.ProfileEditParams{
		Title:   "Phonemeizer - To Phonetics - chuyển sang âm vị",
		Message: "Hello from profile edit",
	}
	view.PhonemizerView(w, p)
}

func (app *Controller) PhonemizerPostCtrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text := strings.TrimSpace(r.FormValue("textarea"))
		language := r.FormValue("language")
		textLower := strings.ToLower(text)

		var message string
		length := len(textLower)

		if length > 2000 {
			fmt.Println(length)
			message = "<p style = 'color:red'>Văn bản của bạn vượt quá 2000 kí tự!</p>"
			fmt.Fprintf(w, "%s", message)
			return
		}

		if length <= 0 {
			message = "<p style = 'color:red'>Bạn chưa nhập văn bản tiếng Anh hoặc tiếng Việt</p>"
			fmt.Fprintf(w, "%s", message)
			return
		}
		if language == "vietnamese" {
			fmt.Fprintf(w, "%s", textLower)
		} else if language == "english" {
			pronunciations, err := featuresmodels.EnglishWord(textLower)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Fprintf(w, "%s", pronunciations)
		} else {
			http.Error(w, "Invalid language selection", http.StatusBadRequest)
		}
		//fmt.Fprintf(w, "%s", "Tính năng này đang được xây dựng!")
	}
}
