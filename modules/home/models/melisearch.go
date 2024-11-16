package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/meilisearch/meilisearch-go"
	"github.com/yeungon/corpora/internal/config"
)

// Item represents the structure of the data
type Item struct {
	Author       string
	Content      string
	CrawledAt    string
	ID           int
	PictureCount int
	Processed    int
	Source       string
	Title        string
	Topic        string
	URL          string
}

// CheckIfKeywordMatches checks if the keyword is present in the long text and returns the match count.
func CheckIfKeywordMatches(text, keyword string) (bool, int) {
	text = strings.ToLower(text)
	keyword = strings.ToLower(keyword)

	// Count occurrences of the keyword
	matches := strings.Count(text, keyword)
	return matches > 0, matches
}

// MeliSearch performs a search on Meilisearch and returns filtered results, total matched count,
// the number of content items that matched, and the total occurrences across all content fields.
func MeliSearch(keyword string) ([]Item, int, int, int, error) {
	searchURL := config.GET().MELISEARCH_URL
	searchKey := config.GET().MELISEARCH_API_KEY
	client := meilisearch.New(searchURL, meilisearch.WithAPIKey(searchKey))

	// Step 1: Update typo tolerance for the "vietnamese_news" index
	_, err := client.Index("vietnamese_news").UpdateTypoTolerance(&meilisearch.TypoTolerance{
		Enabled: false,
		MinWordSizeForTypos: meilisearch.MinWordSizeForTypos{
			OneTypo:  10,
			TwoTypos: 20,
		},
	})
	if err != nil {
		fmt.Println("Error updating typo tolerance:", err)
		os.Exit(1)
	}

	// Step 2: Perform the search
	searchRes, err := client.Index("vietnamese_news").Search(keyword, &meilisearch.SearchRequest{
		AttributesToRetrieve:  []string{"author", "content", "crawled_at", "id", "picture_count", "processed", "source", "title", "topic", "url"},
		AttributesToHighlight: []string{"content"},
		Limit:                 30,
	})
	if err != nil {
		fmt.Println("Search error:", err)
		os.Exit(1)
	}

	// Step 3: Initialize counters
	var filteredResults []Item
	contentMatchedCount := 0 // Number of content items that matched
	totalMatchCount := 0     // Total number of matches across all content fields

	// Step 4: Filter the results based on content
	for _, hit := range searchRes.Hits {
		// Convert the hit to an Item struct
		var item Item
		if content, ok := hit.(map[string]interface{})["content"].(string); ok {
			item.Content = content
		}
		if title, ok := hit.(map[string]interface{})["title"].(string); ok {
			item.Title = title
		}

		// Check if the keyword is present in the content and count occurrences
		matched, matchCount := CheckIfKeywordMatches(item.Content, keyword)
		if matched {
			contentMatchedCount++                           // Increment content match counter
			totalMatchCount += matchCount                   // Increment total match counter by occurrences in this content
			filteredResults = append(filteredResults, item) // Add to results
		}
	}

	// Update totalMatched to be the count of filtered results
	totalMatched := len(filteredResults)

	return filteredResults, totalMatched, contentMatchedCount, totalMatchCount, nil
}
