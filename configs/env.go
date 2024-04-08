package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoURI := os.Getenv("MONGOURI")

	return mongoURI
}
