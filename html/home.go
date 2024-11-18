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

type IndexParams struct {
	Title        string
	Message      string
	StateSearch  bool
	Results      []Item
	CountMatched int
	TotalMatch   int32
	UserData     interface{}
}

func Home(w io.Writer, p IndexParams) error {
	return home.ExecuteTemplate(w, "layout.html", p)
}
