package models

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	manticoreclient "github.com/manticoresoftware/manticoresearch-go"
	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/internal/config"
)

type ManticoreSearchResultMyNews struct {
	Text         string   `json:"text"`
	Domain       string   `json:"domain"`
	Concordances []string `json:"concordances"`
}

func ManticoreMyNews(keyword string, index_selected string, page int) (int32, map[string]interface{}, []html.Concordance) {
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

	// Number of article fetched per page (not the actual concordance. For example, one article might have more than 2 concordances.)
	pageSize := 15
	offset := (page - 1) * pageSize // Calculate offset based on page and pageSize

	// Set limit to 5 results
	searchRequest.SetLimit(int32(pageSize))
	searchRequest.SetOffset(int32(offset))

	// Execute the search request
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		// Return an empty slice and 0 as the total count in case of an error
		return 0, nil, nil
	}

	// Create a slice of ManticoreSearchResult from the response
	var results []ManticoreSearchResultMyNews

	// Total match documents;
	totalMatchedQuery := *resp.Hits.Total
	timeTookToQuery := *resp.Took

	// Iterate through the hits

	for _, hit := range resp.Hits.Hits {
		// Extract the _source field, which is a map
		source := hit["_source"].(map[string]interface{})
		text, text0k := source["text"].(string)
		domain, domainOk := source["domain"].(string)
		// Only append to results if both fields exist and are strings
		if text0k && domainOk {
			results = append(results, ManticoreSearchResultMyNews{
				Text:   text,
				Domain: domain,
			})
		} else {
			// Handle the case where either 'word' or 'define' is missing or not a string
			fmt.Println("Invalid data: word or define missing")
		}
	}
	// left and right text
	window := 15
	// Process each result to extract concordances
	var concordances []html.Concordance

	for _, result := range results {
		extractedConcordances := extractConcordance(result.Text, keyword, window)
		for _, concordance := range extractedConcordances {
			// Split the concordance around the keyword
			splitParts := splitConcordance(concordance, keyword)
			concordances = append(concordances, splitParts)
		}
	}

	total_concordance := len(concordances)

	first50 := getFirstNItems(concordances, 20)

	fmt.Println("tổng số concordance", total_concordance)
	fmt.Println("tổng số totalMatchedQuery", totalMatchedQuery)

	// Calculate totalPages based on total matches and pageSize
	totalPages := (int(totalMatchedQuery) + pageSize - 1) / pageSize // Round up the division

	// Ensure offset does not exceed totalMatches
	if offset >= int(totalMatchedQuery) {
		// Adjust page to the last valid page if offset goes out of range
		page = totalPages
		offset = (totalPages - 1) * pageSize
	}
	// Pagination to pass to HTML template for handling pagination
	pagination := map[string]interface{}{
		"time_took":    timeTookToQuery,
		"page":         page,
		"pageSize":     pageSize,
		"offset":       offset,
		"totalMatches": totalMatchedQuery,
		"totalPages":   totalPages,
		"nextPage":     page + 1,
		"prevPage":     page - 1,
	}

	return totalMatchedQuery, pagination, first50
}

// extractConcordance function with case insensitivity but preserving original case in the output
func extractConcordance(text, phrase string, window int) []string {
	// Convert the text to lowercase for case-insensitive matching
	lowerText := strings.ToLower(text)
	phrase = strings.ToLower(phrase)

	// Tokenize the text into words using a regular expression
	words := regexp.MustCompile(`\S+`).FindAllString(text, -1)
	lowerWords := regexp.MustCompile(`\S+`).FindAllString(lowerText, -1)

	phraseWords := strings.Fields(phrase) // Split the phrase into individual words
	phraseLen := len(phraseWords)

	var concordances []string

	// Loop through the lowercase words to find the phrase match
	for i := 0; i <= len(lowerWords)-phraseLen; i++ {
		// Check if the current slice of lowercase words matches the phrase (case-insensitive)
		if strings.Join(lowerWords[i:i+phraseLen], " ") == phrase {
			// Calculate start and end indices for the window
			start := max(0, i-window)
			end := min(len(words), i+phraseLen+window)

			// Extract the concordance slice based on the original words (preserving the case)
			concordance := strings.Join(words[start:end], " ")
			concordances = append(concordances, concordance)
		}
	}
	return concordances
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitConcordance(concordance, keyword string) html.Concordance {
	parts := html.Concordance{}

	// Convert the concordance and keyword to lowercase for case-insensitive matching
	lowerConcordance := strings.ToLower(concordance)
	lowerKeyword := strings.ToLower(keyword)

	// Find the index of the lowercase keyword in the lowercase concordance
	index := strings.Index(lowerConcordance, lowerKeyword)
	if index != -1 {
		// Map the lowercase match to the original case in the concordance
		parts.BeforeKeyword = concordance[:index]
		parts.Keyword = concordance[index : index+len(keyword)]
		parts.AfterKeyword = concordance[index+len(keyword):]
	} else {
		parts.BeforeKeyword = concordance
		parts.Keyword = "" // No keyword found
		parts.AfterKeyword = ""
	}

	return parts
}

func getFirstNItems(slice []html.Concordance, n int) []html.Concordance {
	if len(slice) > n {
		return slice[:n]
	}
	return slice
}
