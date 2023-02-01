package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dbUser := GetEnvValue("DB_USER")
	log.Println(dbUser)
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
