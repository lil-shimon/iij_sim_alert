package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("failed to read .env: %v", err)
	}

	tkn := os.Getenv("TEST_API_TOKEN")
	url := os.Getenv("TEST_CHAT_URL")

	return tkn, url
}
