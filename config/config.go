package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppName string
	AppHost string
	AppPort string
}

var appConfig *Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig = &Config{
		AppName: GetString("APP_NAME"),
		AppHost: GetString("APP_HOST"),
		AppPort: GetString("APP_PORT"),
	}
}

func GetAppName() string {
	return appConfig.AppName
}

func GetAppHost() string {
	return appConfig.AppHost
}

func GetAppPort() string {
	return appConfig.AppPort
}

func GetString(key string) string {
	return os.Getenv(key)
}
