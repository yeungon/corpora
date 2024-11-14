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
	data, err := models.MeliSearch(query)

	// Convert the []interface{} into JSON []byte
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	var items []html.Item
	err = json.Unmarshal([]byte(dataBytes), &items)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	p := html.IndexParams{
		Title:       "Vietnamese Corpora",
		Message:     query,
		StateSearch: true,
		Results:     items,
	}

	html.Home(w, p)
}
