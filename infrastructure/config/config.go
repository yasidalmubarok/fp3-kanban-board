package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBDialect     string
	Port          string
	SecretKey     string
	AdminFullName string
	AdminEmail    string
	AdminPassword string
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("error while processing .env file")
	}
}

func AppConfig() appConfig {
	return appConfig{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBDialect:     os.Getenv("DB_DIALECT"),
		Port:          os.Getenv("PORT"),
		SecretKey:     os.Getenv("SECRET_KEY"),
		AdminFullName: os.Getenv("ADMIN_FULL_NAME"),
		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}
}
