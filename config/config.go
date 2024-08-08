package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtKey []byte

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JwtKey = []byte(os.Getenv("JWT_KEY"))
}