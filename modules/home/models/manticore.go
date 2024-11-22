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

type ManticoreSearchResult struct {
	Word   string `json:"word"`
	Define string `json:"define"`
}

type ManticoreSearchResultMyNews struct {
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Concordances []string `json:"concordances"`
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

	// Set limit to 5 results
	searchRequest.SetLimit(50)

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

func ManticoreMyNews(keyword string, index_selected string, page int) ([]ManticoreSearchResultMyNews, int32, map[string]interface{}) {
	configuration := manticoreclient.NewConfiguration()
	searchURL := config.GET().MANTICORESEARCH_URL
	configuration.Servers[0].URL = searchURL
	apiClient := manticoreclient.NewAPIClient(configuration)
	searchRequest := *manticoreclient.NewSearchRequest(index_selected)
	// Option 1: Also return RELAVANT words/phrase
	// query := map[string]interface{}{
	// 	"match": map[string]interface{}{
	// 		"*": keyword,
	// 	},
	// }

	// Option 2: Onlyreturn matched words/phrase
	query := map[string]interface{}{
		"match_phrase": map[string]interface{}{
			"*": keyword, // Matches the entire phrase across all fields
		},
	}

	searchRequest.SetQuery(query)

	// highlight := manticoreclient.NewHighlight()
	// searchRequest.SetHighlight(*highlight)
	// searchRequest.Options = options

	// Define page and pageSize for pagination
	pageSize := 20                  // Number of results per page
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
		return []ManticoreSearchResultMyNews{}, 0, nil
	}

	// Create a slice of ManticoreSearchResult from the response
	var results []ManticoreSearchResultMyNews

	// Total match documents;
	totalMatchedQuery := *resp.Hits.Total

	fmt.Println("totalMatchedQuery", totalMatchedQuery)

	timeTookToQuery := *resp.Took

	// Iterate through the hits

	for _, hit := range resp.Hits.Hits {
		// Extract the _source field, which is a map
		source := hit["_source"].(map[string]interface{})
		title, titleOk := source["title"].(string)
		content, contentOk := source["content"].(string)
		// Only append to results if both fields exist and are strings
		if titleOk && contentOk {
			results = append(results, ManticoreSearchResultMyNews{
				Title:   title,
				Content: content,
			})
		} else {
			// Handle the case where either 'word' or 'define' is missing or not a string
			fmt.Println("Invalid data: word or define missing")
		}
	}

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

	// window := 10

	// // Process each result to extract concordances
	// var concordances []html.Concordance
	// for _, result := range results {
	// 	extractedConcordances := extractConcordance(result.Content, keyword, window)
	// 	for _, concordance := range extractedConcordances {
	// 		// Split the concordance around the keyword
	// 		splitParts := splitConcordance(concordance, keyword)
	// 		concordances = append(concordances, splitParts)
	// 	}
	// }

	// Return the results and the total
	return results, totalMatchedQuery, pagination
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
