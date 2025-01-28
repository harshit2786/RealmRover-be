package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost  string
	DBUser  string
	DBPass  string
	DBName  string
	DBPort  string
	Port    string
	JWT     string
	GITPASS string
	GITCLI  string
	REDURI  string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file. Using default values.")
	}
	return &Config{
		DBHost:  getEnv("HOST", "localhost"),
		DBUser:  getEnv("DBUSER", "postgres"),
		DBPass:  getEnv("DBPASS", "mysecretpassword"),
		DBName:  getEnv("DBNAME", "postgres"),
		DBPort:  getEnv("DBPORT", "5432"),
		Port:    getEnv("PORT", "8000"),
		JWT:     getEnv("JWT_SECRET", "2891142ba5c97f516a42"),
		GITPASS: getEnv("GITHUB_SECRET", ""),
		GITCLI: getEnv("GITHUB_CLIENT",""),
		REDURI: getEnv("REDIRECT_URL","http://localhost:5173/callback/github"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return defaultValue
	}
}
