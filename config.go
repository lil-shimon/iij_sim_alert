package main

import (
	"fmt"
	"os"
)

func loadEnv() (string, string) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("failed to read .env: %v", err)
	}

	tkn := os.Getenv("API_TOKEN")
	url := os.Getenv("CHAT_URL")

	return tkn, url
}