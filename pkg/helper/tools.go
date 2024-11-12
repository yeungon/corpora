package helper

import (
	"fmt"
	"os"
)

func createDirectory(path string) error {
	err := os.MkdirAll(path, 0755) // 0755 gives read/write/execute permissions to the owner and read/execute permissions to others
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}
