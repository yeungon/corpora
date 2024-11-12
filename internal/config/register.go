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
	TEST string
}

// Register the config
func New() *ENV {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		test_env := os.Getenv("TEST")
		env = &ENV{
			TEST: test_env,
		}
	})
	return env
}

func GET() *ENV {
	return env
}
