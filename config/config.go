package config

import (
	"os"
)

var JwtKey []byte

func LoadConfig() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	JwtKey = []byte(os.Getenv("JWT_KEY"))
}