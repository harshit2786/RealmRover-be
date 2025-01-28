package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
	Port   string
}

func GetConfig() *Config {
	err := godotenv.Load()
    if err != nil {
        log.Println("Warning: Could not load .env file. Using default values.")
    }
	return &Config{
		DBHost: getEnv("HOST","localhost"),
		DBUser: getEnv("DBUSER","postgres"),
		DBPass: getEnv("DBPASS","mysecretpassword"),
		DBName: getEnv("DBNAME","postgres"),
		DBPort: getEnv("DBPORT","5432"),
		Port: getEnv("PORT","8000"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if(exists){
		return value
	} else {
		return defaultValue
	}
}