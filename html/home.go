package html

import (
	"io"
)

// Define the structure for a single item
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

type IndexParams struct {
	Title       string
	Message     string
	StateSearch bool
	Results     []Item
}

func Home(w io.Writer, p IndexParams) error {
	return home.ExecuteTemplate(w, "layout.html", p)
}
