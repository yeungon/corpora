package models

import (
	"context"
	"fmt"
	"os"

	manticoreclient "github.com/manticoresoftware/manticoresearch-go"
	"github.com/yeungon/corpora/internal/config"
)

type ManticoreSearchResult struct {
	Word   string `json:"word"`
	Define string `json:"define"`
}

func ManticoreDictionary(keyword string, index_selected string) ([]ManticoreSearchResult, int32) {
	configuration := manticoreclient.NewConfiguration()
	searchURL := config.GET().MANTICORESEARCH_URL
	configuration.Servers[0].URL = searchURL
	apiClient := manticoreclient.NewAPIClient(configuration)

	searchRequest := *manticoreclient.NewSearchRequest(index_selected)
	// Option 2: Onlyreturn matched words/phrase
	query := map[string]interface{}{
		"match_phrase": map[string]interface{}{
			"*": keyword, // Matches the entire phrase across all fields
		},
	}

	searchRequest.SetQuery(query)

	// Set limit to 50 results
	searchRequest.SetLimit(1000)

	// Execute the search request
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		// Return an empty slice and 0 as the total count in case of an error
		return []ManticoreSearchResult{}, 0
	}

	// Create a slice of ManticoreSearchResult from the response
	var results []ManticoreSearchResult

	// Get the total hits
	total := *resp.Hits.Total
	fmt.Println("Index_lsselected", index_selected)
	fmt.Println("Total hits:", total)

	// Iterate through the hits
	for _, hit := range resp.Hits.Hits {
		// Extract the _source field, which is a map
		source := hit["_source"].(map[string]interface{})
		// Extract 'word' and 'define' from the source map
		word, wordOk := source["word"].(string)
		define, defineOk := source["define"].(string)
		// Only append to results if both fields exist and are strings
		if wordOk && defineOk {
			results = append(results, ManticoreSearchResult{
				Word:   word,
				Define: define,
			})
		} else {
			// Handle the case where either 'word' or 'define' is missing or not a string
			fmt.Println("Invalid data: word or define missing")
		}
	}

	// Return the results and the total
	return results, total
}
