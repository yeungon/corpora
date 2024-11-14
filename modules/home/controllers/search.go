package home

import (
	"encoding/json"
	"fmt"
	"net/http"

	html "github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/modules/home/models"
)

func (app *Controller) SearchConcordancePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	query := r.URL.Query().Get("keyword")
	selectedOption := r.URL.Query().Get("corpusOptions")

	fmt.Println("Selected option:", selectedOption) // Perform search and retrieve results
	results, err := models.MeliSearch(query)
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	// Convert results to JSON string for display
	resultsJSON, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}

	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     query,
		StateSearch: true,
		Results:     string(resultsJSON),
	}

	html.Home(w, p)
}
