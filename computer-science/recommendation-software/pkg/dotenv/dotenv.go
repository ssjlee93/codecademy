package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
