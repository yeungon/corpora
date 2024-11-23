package html

import (
	"io"
)

// Define the structure for a single item
type Item struct {
	ID           int `json:"_id"`
	Score        int `json:"_score"`
	Author       string
	Content      string
	CrawledAt    string
	PictureCount int
	Processed    int
	Source       string
	Title        string
	Topic        string
	URL          string
	Define       string `json:"_source.define"`
	Word         string `json:"_source.word"`
}

// Define a struct to hold the parts of the concordance
type Concordance struct {
	BeforeKeyword string // Text before the keyword
	Keyword       string // The keyword itself
	AfterKeyword  string // Text after the keyword
}
type IndexParams struct {
	Title        string
	Message      string
	SourceIndex  string
	StateSearch  bool
	Results      []Item
	Concordances []Concordance
	CountMatched int
	TotalMatch   int32
	UserData     interface{}
	CurrentURL   string
	Page         int
	Pagination   map[string]interface{}
}

func Home(w io.Writer, p IndexParams) error {
	return home.ExecuteTemplate(w, "layout.html", p)
}

func Credit(w io.Writer, p IndexParams) error {
	return pageCredit.ExecuteTemplate(w, "layout.html", p)
}
