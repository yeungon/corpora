package featuresmodels

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func EnglishWord(paragraph string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}

	data, err := os.ReadFile(dir + "/privatedata/compress_open_ipa_en_US.json")
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Parse JSON into a map
	var content map[string][]map[string]string
	if err := json.Unmarshal(data, &content); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Build a pronunciations map
	pronunciations := make(map[string]string)
	if enUS, ok := content["en_US"]; ok {
		for _, entry := range enUS {
			for word, pronunciation := range entry {
				pronunciations[word] = pronunciation
			}
		}
	}

	// Process the paragraph to maintain order
	var pronunciationParts []string
	words := strings.Fields(paragraph) // Split the paragraph into words
	for _, word := range words {
		normalizedWord := strings.Trim(word, ",.?!'\"") // Normalize the word
		if pronunciation, exists := pronunciations[normalizedWord]; exists && pronunciation != "Not found" {
			// Process the pronunciation to replace inner slashes with "|"
			cleanedPronunciation := simplifyPronunciation(pronunciation)
			pronunciationParts = append(pronunciationParts, cleanedPronunciation)
		} else {
			pronunciationParts = append(pronunciationParts, normalizedWord) // Include the original word if not found
		}
	}

	// Join all parts and wrap in "/"
	result := "/" + strings.Join(pronunciationParts, " ") + "/"
	return result, nil
}

// simplifyPronunciation replaces inner slashes with "|" correctly
func simplifyPronunciation(pronunciation string) string {
	if len(pronunciation) > 2 && pronunciation[0] == '/' && pronunciation[len(pronunciation)-1] == '/' {
		content := pronunciation[1 : len(pronunciation)-1] // Remove the first and last slashes
		// Replace ", " (comma and space) with "|"
		cleaned := strings.ReplaceAll(content, ", ", "|")
		return cleaned
	}
	return pronunciation
}
