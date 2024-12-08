package featuresmodels

import (
	"fmt"
	"os"
)

func VietnameseWord() []byte {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	data, err := os.ReadFile(dir + "/privatedata/vietnameseword.txt")
	if err != nil {
		panic(err)
	}
	return data
}
