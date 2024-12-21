package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port string
	}
	Database Database
}

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

var appConfig Config

func LoadConfig() {
	errorMessage := godotenv.Load()
	if errorMessage != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.Server.Port = ":" + os.Getenv("SERVER_PORT")
	appConfig.Database.Host = os.Getenv("DB_HOST")
	appConfig.Database.User = os.Getenv("DB_USER")
	appConfig.Database.Password = os.Getenv("DB_PASSWORD")
	appConfig.Database.Name = os.Getenv("DB_NAME")
	appConfig.Database.Port, errorMessage = strconv.Atoi(os.Getenv("DB_PORT"))

	if errorMessage != nil {
		log.Fatalf("Error loading .env file: %s", errorMessage)
	}
}

func GetConfig() *Config {
	return &appConfig
}
