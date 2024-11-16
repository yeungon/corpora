package models

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	manticoreclient "github.com/manticoresoftware/manticoresearch-go"
	"github.com/yeungon/corpora/internal/config"
)

func Manticore() {
	configuration := manticoreclient.NewConfiguration()
	searchURL := config.GET().MANTICORESEARCH_URL
	configuration.Servers[0].URL = searchURL
	apiClient := manticoreclient.NewAPIClient(configuration)

	// Prepare a search request for the "my_index" index
	searchRequest := *manticoreclient.NewSearchRequest("my_index")

	// Create a match query similar to the one in your curl command
	query := map[string]interface{}{
		"match": map[string]interface{}{
			"*": "learn",
		},
	}
	searchRequest.SetQuery(query)

	// Execute the search request
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	// Marshal the response to JSON and print it
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling response: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`:\n%s\n", jsonData)
}
