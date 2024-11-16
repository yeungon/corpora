package home

import (
	"fmt"
	"net/http"

	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/modules/home/models"
)

// Define a struct for Manticore search results to unmarshal into
type ManticoreSearchResult struct {
	ID     int    `json:"_id"`
	Score  int    `json:"_score"`
	Define string `json:"_source.define"`
	Word   string `json:"_source.word"`
}

func (app *Controller) SearchManticore(w http.ResponseWriter, r *http.Request) {

	// Max body size for request
	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	// Get the search keyword from the URL query parameter
	query := r.URL.Query().Get("keyword")
	fmt.Printf("Received search query: %s\n", query)

	// Call the Manticore model to get search results
	searchResults, total := models.Manticore(query)

	_ = searchResults

	// Initialize a slice to hold the parsed items
	var items []html.Item

	// Convert the Manticore search results into html.Item for rendering
	// for _, result := range searchResults {
	// 	items = append(items, html.Item{
	// 		Word:   result.Word,
	// 		Define: result.Define,
	// 	})
	// }

	// Prepare the IndexParams for the HTML page
	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     query,
		StateSearch: true,
		Results:     items,
		TotalMatch:  total,
	}

	// Render the Home page template with the search results
	html.Home(w, p)
}
