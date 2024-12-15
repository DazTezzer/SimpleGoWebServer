package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDBHost() string {
	return os.Getenv("DATABASE_HOST")
}

func GetDBPort() string {
	return os.Getenv("DATABASE_PORT")
}

func GetDBUser() string {
	return os.Getenv("DATABASE_USER")
}

func GetDBPassword() string {
	return os.Getenv("DATABASE_PASSWORD")
}

func GetDBName() string {
	return os.Getenv("DATABASE_NAME")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetCORS() string {
	return os.Getenv("CORS")
}
