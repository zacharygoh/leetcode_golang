package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Environment Checking
	if os.Getenv("ENV") == "development" {
		Host = os.Getenv("DEV_DB_HOST")
		Port = os.Getenv("DEV_DB_PORT")
		User = os.Getenv("DEV_DB_USER")
		Password = os.Getenv("DEV_DB_PASSWORD")
		Dbname = os.Getenv("DEV_DB_NAME")
		Url = os.Getenv("DEV_URL")
	} else if os.Getenv("ENV") == "testing" {
		Host = os.Getenv("TEST_DB_HOST")
		Port = os.Getenv("TEST_DB_PORT")
		User = os.Getenv("TEST_DB_USER")
		Password = os.Getenv("TEST_DB_PASSWORD")
		Dbname = os.Getenv("TEST_DB_NAME")
		Url = os.Getenv("TEST_URL")
	} else {
		Host = os.Getenv("PROD_DB_HOST")
		Port = os.Getenv("PROD_DB_PORT")
		User = os.Getenv("PROD_DB_USER")
		Password = os.Getenv("PROD_DB_PASSWORD")
		Dbname = os.Getenv("PROD_DB_NAME")
		Url = os.Getenv("PROD_URL")
	}

}
