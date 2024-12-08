package featuresmodels

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// EnglishWord processes a paragraph and returns a map of words to their pronunciations.
func EnglishWord(paragraph string) (map[string]string, error) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Build the file path
	filePath := filepath.Join(dir, "privatedata/compress_open_ipa_en_US.json")

	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse JSON into a map
	var content map[string][]map[string]string
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Split the paragraph into words
	words := strings.Fields(paragraph)

	// Prepare the result map
	pronunciations := make(map[string]string)

	// Loop through the JSON structure
	if enUS, ok := content["en_US"]; ok {
		for _, word := range words {
			normalizedWord := strings.Trim(word, ",.?!'\"") // Normalize the word (remove punctuation)
			for _, entry := range enUS {
				if value, exists := entry[normalizedWord]; exists {
					pronunciations[normalizedWord] = value
					break
				}
			}
			// If not found, indicate it
			if _, found := pronunciations[normalizedWord]; !found {
				pronunciations[normalizedWord] = "Not found"
			}
		}
	}

	return pronunciations, nil
}
