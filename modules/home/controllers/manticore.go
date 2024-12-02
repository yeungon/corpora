package home

import (
	"fmt"
	"net/http"
	"net/url"
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
var concordances []html.Concordance
var concordances_model []html.Concordance

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
	index_selected := "my_index"
	if index_selected == "my_news" {
		source = "english"
		items, total = SearchEnglish(query, index_selected)

	}
	if index_selected == "my_index" {
		source = "vietnamese_news"
		total, pagination, concordances_model = SearchMyNews(query, index_selected, page)

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
		Concordances: concordances_model,
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

func SearchMyNews(query string, index_selected string, page int) (int32, map[string]interface{}, []html.Concordance) {
	total, pagination, concordances := models.ManticoreMyNews(query, index_selected, page)
	return total, pagination, concordances
}
