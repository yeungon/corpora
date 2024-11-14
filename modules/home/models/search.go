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
	// Meilisearch is typo-tolerant:
	searchRes, err := client.Index("vietnamese_news").Search(keyword,
		&meilisearch.SearchRequest{
			Limit: 50,
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return searchRes.Hits, nil
}
