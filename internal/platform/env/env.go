package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	JWTSecret    string
	DatabaseUrl  string
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		DatabaseUrl:  os.Getenv("DATABASE_URL"),	
	}
}