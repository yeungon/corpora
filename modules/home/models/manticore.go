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

func Manticore(keyword string) ([]ManticoreSearchResult, int32) {
	configuration := manticoreclient.NewConfiguration()
	searchURL := config.GET().MANTICORESEARCH_URL
	configuration.Servers[0].URL = searchURL
	apiClient := manticoreclient.NewAPIClient(configuration)

	// Prepare a search request for the "my_index" index
	searchRequest := *manticoreclient.NewSearchRequest("my_index")

	query := map[string]interface{}{
		"match": map[string]interface{}{
			"*": keyword,
		},
	}
	searchRequest.SetQuery(query)

	// Set limit to 5 results
	searchRequest.SetLimit(5)

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

	total := *resp.Hits.Total
	fmt.Println(total)

	// Iterate through the hits and fill in the results slice
	// for _, hit := range resp.Hits.Hits {
	// 	results = append(results, ManticoreSearchResult{
	// 		Word:   hit["word"].(string),
	// 		Define: hit["define"].(string),
	// 	})
	// }

	// Return the results and the total
	return results, total
}
