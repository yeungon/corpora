package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once //
	env  *ENV      //
)

type ENV struct {
	APP                 string
	APPURL              string
	TEST                string
	MELISEARCH_URL      string
	MELISEARCH_API_KEY  string
	MANTICORESEARCH_URL string
}

// Register the config
func New() *ENV {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		app := os.Getenv("APP")
		app_url := os.Getenv("APPURL")
		test_env := os.Getenv("TEST")
		search_url := os.Getenv("MELISEARCH_URL")
		search_api := os.Getenv("MELISEARCH_API_KEY")
		manticore_url := os.Getenv("MANTICORESEARCH_URL")

		env = &ENV{
			TEST:                test_env,
			MELISEARCH_URL:      search_url,
			MELISEARCH_API_KEY:  search_api,
			MANTICORESEARCH_URL: manticore_url,
			APPURL:              app_url,
			APP:                 app,
		}
	})
	return env
}

func GET() *ENV {
	return env
}
