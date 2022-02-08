package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// EnvMongoURI checks if env variable loaded correctly & return env variable
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}