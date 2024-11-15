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
	search_url := config.GET().MANTICORESEARCH_URL
	configuration.Servers[0].URL = search_url
	apiClient := manticoreclient.NewAPIClient(configuration)

	// // Add documents to an index
	// docs := []string{
	// 	"{\"insert\": {\"index\" : \"test\", \"id\" : 1, \"doc\" : {\"title\" : \"Title 1\"}}}",
	// 	"{\"insert\": {\"index\" : \"test\", \"id\" : 2, \"doc\" : {\"title\" : \"Title 2\"}}}",
	// }
	// apiClient.IndexAPI.Bulk(context.Background()).Body(strings.Join(docs[:], "\n")).Execute()

	// response from `Search`: SearchRequest
	searchRequest := *manticoreclient.NewSearchRequest("test")
	// Perform a search
	query := map[string]interface{}{"query_string": "Title"}
	searchRequest.SetQuery(query)
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Marshal the response to JSON
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling response: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`:\n%s\n", jsonData)

}
