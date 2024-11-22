package home

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/yeungon/corpora/html"
	"github.com/yeungon/corpora/internal/config"
	"github.com/yeungon/corpora/modules/home/models"
)

type SearchData struct {
	Keyword       string
	CorpusOptions string
	Source        string
}

var items []html.Item
var total int32
var source string
var pagination map[string]interface{}

func (app *Controller) SearchManticore(w http.ResponseWriter, r *http.Request) {
	// Form Data
	query := r.URL.Query().Get("keyword")
	query = strings.TrimSpace(query)

	if len(query) <= 0 {
		var something = "wroigggg"
		fmt.Printf("Something %s", something)

	}

	selectedOption := r.URL.Query().Get("corpusOptions")

	// URL Data
	queryParams := r.URL.Query()
	pageParams := queryParams["page"]
	baseURL := config.GET().APPURL
	fullURL := fmt.Sprintf("%s%s", baseURL, r.RequestURI)

	// Parse the URL
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Get the query parameters
	queryParamsUpdate := parsedURL.Query()
	// Remove the "page" parameter
	queryParamsUpdate.Del("page")
	// Set the modified query parameters back to the URL
	parsedURL.RawQuery = queryParamsUpdate.Encode()
	// Get the updated URL as a string
	updatedURL := parsedURL.String()

	var page int
	if len(pageParams) > 0 {
		// Convert the first element to an int
		parsedPage, err := strconv.Atoi(pageParams[0])
		if err != nil {
			fmt.Println("Error parsing page:", err)
			return
		}
		page = parsedPage
	} else {
		page = 1
	}

	//index_selected := "poetic_nom"
	index_selected := "my_news"
	if index_selected == "my_index" {
		source = "english"
		items, total = SearchEnglish(query, index_selected)

	}
	if index_selected == "my_news" {
		source = "vietnamese_news"
		items, total, pagination = SearchMyNews(query, index_selected, page)

	}

	keyword := query
	window := 20

	// Process each result to extract concordances
	var concordances []html.Concordance
	for _, result := range items {
		extractedConcordances := extractConcordance(result.Content, keyword, window)
		for _, concordance := range extractedConcordances {
			// Split the concordance around the keyword
			splitParts := splitConcordance(concordance, keyword)
			concordances = append(concordances, splitParts)
		}
	}

	SearchDataInstance := SearchData{
		Keyword:       query,
		CorpusOptions: selectedOption,
		Source:        source,
	}
	// Prepare the IndexParams for the HTML page
	p := html.IndexParams{
		Title:        "Vietnamese Corpora",
		Message:      query,
		SourceIndex:  source,
		StateSearch:  true,
		Results:      items,
		Concordances: concordances,
		TotalMatch:   total,
		UserData:     SearchDataInstance,
		CurrentURL:   updatedURL,
		Page:         page,
		Pagination:   pagination,
	}

	// Render the Home page template with the search results
	html.Home(w, p)
}

func SearchEnglish(query string, index_selected string) ([]html.Item, int32) {
	var items []html.Item
	searchResults, total := models.ManticoreDictionary(query, index_selected)
	for _, result := range searchResults {
		items = append(items, html.Item{
			Word:   result.Word,
			Define: result.Define,
		})
	}
	return items, total

}

func SearchMyNews(query string, index_selected string, page int) ([]html.Item, int32, map[string]interface{}) {
	var items []html.Item
	searchResults, total, pagination := models.ManticoreMyNews(query, index_selected, page)
	for _, result := range searchResults {
		items = append(items, html.Item{
			Title:   result.Title,
			Content: result.Content,
		})
	}
	return items, total, pagination

}

//	type Item struct {
//		ID           int `json:"_id"`
//		Score        int `json:"_score"`
//		Author       string
//		Content      string
//		CrawledAt    string
//		PictureCount int
//		Processed    int
//		Source       string
//		Title        string
//		Topic        string
//		URL          string
//		Define       string `json:"_source.define"`
//		Word         string `json:"_source.word"`
//	}
//
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
