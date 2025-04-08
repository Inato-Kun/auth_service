package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env cannot to load %v", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
