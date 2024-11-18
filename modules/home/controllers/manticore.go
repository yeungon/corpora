package home

import (
	"net/http"

	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/modules/home/models"
)

type SearchData struct {
	Keyword       string
	CorpusOptions string
	Source        string
}

var items []html.Item
var total int32
var source string

func (app *Controller) SearchManticore(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("keyword")
	selectedOption := r.URL.Query().Get("corpusOptions")
	//index_selected := "poetic_nom"
	index_selected := "my_news"

	if index_selected == "my_index" {
		source = "english"
		items, total = SearchEnglish(query, index_selected)

	}

	if index_selected == "my_news" {
		source = "vietnamese_news"
		items, total = SearchMyNews(query, index_selected)

	}

	SearchDataInstance := SearchData{
		Keyword:       query,
		CorpusOptions: selectedOption,
		Source:        source,
	}
	// Prepare the IndexParams for the HTML page
	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     query,
		SourceIndex: source,
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

func SearchMyNews(query string, index_selected string) ([]html.Item, int32) {
	var items []html.Item
	searchResults, total := models.ManticoreMyNews(query, index_selected)
	for _, result := range searchResults {
		items = append(items, html.Item{
			Title:   result.Title,
			Content: result.Content,
		})
	}
	return items, total

}

// type Item struct {
// 	ID           int `json:"_id"`
// 	Score        int `json:"_score"`
// 	Author       string
// 	Content      string
// 	CrawledAt    string
// 	PictureCount int
// 	Processed    int
// 	Source       string
// 	Title        string
// 	Topic        string
// 	URL          string
// 	Define       string `json:"_source.define"`
// 	Word         string `json:"_source.word"`
// }
