package home

import (
	"net/http"

	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/modules/home/models"
)

type SearchData struct {
	Keyword       string
	CorpusOptions string
}

var items []html.Item
var total int32

func (app *Controller) SearchManticore(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("keyword")
	selectedOption := r.URL.Query().Get("corpusOptions")
	SearchDataInstance := SearchData{
		Keyword:       query,
		CorpusOptions: selectedOption,
	}

	//index_selected := "poetic_nom"
	index_selected := "my_index"

	if index_selected == "my_index" {
		items, total = SearchEnglish(query, index_selected)

	}

	// Prepare the IndexParams for the HTML page
	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     query,
		StateSearch: true,
		Results:     items,
		TotalMatch:  total,
		UserData:    SearchDataInstance,
	}

	// Render the Home page template with the search results
	html.Home(w, p)
}

func SearchEnglish(query string, index_selected string) ([]html.Item, int32) {
	var items []html.Item
	searchResults, total := models.Manticore(query, index_selected)
	for _, result := range searchResults {
		items = append(items, html.Item{
			Word:   result.Word,
			Define: result.Define,
		})
	}
	return items, total

}
