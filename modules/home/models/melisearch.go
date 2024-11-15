package models

import (
	"fmt"
	"os"

	"github.com/meilisearch/meilisearch-go"
	"github.com/yeungon/corpora/internal/config"
)

func MeliSearch(keyword string) ([]interface{}, error) {
	search_url := config.GET().MELISEARCH_URL
	search_key := config.GET().MELISEARCH_API_KEY
	client := meilisearch.New(search_url, meilisearch.WithAPIKey(search_key))

	// Step 1: Update typo tolerance for the "vietnamese_news" index
	//https://www.meilisearch.com/docs/learn/relevancy/typo_tolerance_calculations
	_, err := client.Index("vietnamese_news").UpdateTypoTolerance(&meilisearch.TypoTolerance{
		Enabled: false,
		MinWordSizeForTypos: meilisearch.MinWordSizeForTypos{
			OneTypo:  20,
			TwoTypos: 30,
		},
	})
	if err != nil {
		fmt.Println("Error updating typo tolerance:", err)
		os.Exit(1)
	}

	searchRes, err := client.Index("vietnamese_news").Search(keyword,
		&meilisearch.SearchRequest{
			AttributesToRetrieve:  []string{"content", "line_number", "title"}, // Specify fields to retrieve
			AttributesToHighlight: []string{"content"},
			HighlightPreTag:       "<span class=\"highlight\">",
			HighlightPostTag:      "</span>", // Highlights matching text within lines
			Limit:                 50,
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return searchRes.Hits, nil
}
