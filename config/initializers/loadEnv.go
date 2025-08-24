package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		_ = godotenv.Load()
	}
}
