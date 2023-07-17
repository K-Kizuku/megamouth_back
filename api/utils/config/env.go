package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiPort       string
	GRPCAddress   string
	DbHost        string
	DbUser        string
	DbPass        string
	DbName        string
	DbPort        string
	TokenLifeTime string
	SigningKey    []byte
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(".env is not found")
	}
	ApiPort = getEnv("API_PORT", "8000")
	GRPCAddress = getEnv("GRPC_ADDRESS", "50052")

	DbHost = getEnv("POSTGRES_HOST", "localhost")
	DbUser = getEnv("POSTGRES_USER", "user")
	DbPass = getEnv("POSTGRES_PASSWORD", "password")
	DbName = getEnv("POSTGRES_DB", "postgres")
	DbPort = getEnv("POSTGRES_PORT", "5432")
	TokenLifeTime = getEnv("TOKEN_LIFETIME", "3600")
	secretKey := getEnv("SECRET_KEY", "secret")
	SigningKey = []byte(secretKey)
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
