package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yeungon/corpora/html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := html.IndexParams{
		Title:   "Vietnamese Corpora",
		Message: "Hello from Index \n",
	}
	html.Index(w, p, Partial(r))
}

func ProfileShow(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}

	// Example usage
	dirPath := "./golang"
	err := createDirectory(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Directory created successfully:", dirPath)
	}

	html.ProfileShow(w, p, Partial(r))
}

func ProfileEdit(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	html.ProfileEdit(w, p, Partial(r))
}

func createDirectory(path string) error {
	err := os.MkdirAll(path, 0755) // 0755 gives read/write/execute permissions to the owner and read/execute permissions to others
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func Partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
