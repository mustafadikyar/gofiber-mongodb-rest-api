package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetConnectionString .env dosyasından MongoDB bağlantı dizesini döndürür.
func GetConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoDbConnectionString := os.Getenv("ConnectionString")
	if mongoDbConnectionString == "" {
		log.Fatal("ConnectionString not found in .env file")
	}

	return mongoDbConnectionString
}
