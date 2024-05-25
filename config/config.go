package config

import (
	"os"

	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/ports"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV               string
	SECRET_KEY        string
	SERVER_PORT       string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	SERVICE_TABLE     string
	REVIEWS_TABLE     string
	REQUEST_TABLE     string
	DEBUG             bool
	TEST              bool
}

func NewConfig(logger ports.LoggerService) (*Config, error) {
	ENV := os.Getenv("ENV")

	switch ENV {
	case "development":
		err := godotenv.Load(".env")
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

	}

	var (
		SECRET_KEY        = os.Getenv("SECRET_KEY")
		SERVER_PORT       = "5001"
		POSTGRES_DB       = "usafihub-cleaner-service"
		POSTGRES_HOST     = "postgres"
		POSTGRES_PORT     = "5432"
		POSTGRES_USER     = "postgres"
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		SERVICE_TABLE     = ""
		REVIEWS_TABLE     = ""
		REQUEST_TABLE     = ""
		DEBUG             = false
		TEST              = false
	)

	switch ENV {
	case "production":
		TEST = false
		DEBUG = false

	case "production_test":
		TEST = true
		DEBUG = true
		SERVICE_TABLE = "Prod_Test_Service"
		REVIEWS_TABLE = "Prod_Test_Review"
		REQUEST_TABLE = "Prod_Test_request"

	case "development":
		TEST = true
		DEBUG = true
		POSTGRES_HOST = "localhost"
		SERVICE_TABLE = "Dev_Service"
		REVIEWS_TABLE = "Dev_Review"
		REQUEST_TABLE = "Dev_request"

	case "development_test":
		TEST = true
		DEBUG = true
		SECRET_KEY = "testsecret"
		POSTGRES_PASSWORD = "pass1234"
		POSTGRES_HOST = "localhost"
		SERVICE_TABLE = "Test_Dev_Service"
		REVIEWS_TABLE = "Test_Dev_Review"
		REQUEST_TABLE = "Test_Dev_request"

	case "docker":
		TEST = true
		DEBUG = true
		SERVICE_TABLE = "Docker_Service"
		REVIEWS_TABLE = "Docker_Review"
		REQUEST_TABLE = "Docker_request"

	case "docker_test":
		TEST = true
		DEBUG = true
		SERVICE_TABLE = "Test_Docker_Service"
		REVIEWS_TABLE = "Test_Docker_Review"
		REQUEST_TABLE = "Test_Docker_request"
	}

	config := Config{
		ENV:               ENV,
		SECRET_KEY:        SECRET_KEY,
		SERVER_PORT:       SERVER_PORT,
		POSTGRES_DB:       POSTGRES_DB,
		POSTGRES_HOST:     POSTGRES_HOST,
		POSTGRES_PORT:     POSTGRES_PORT,
		POSTGRES_USER:     POSTGRES_USER,
		POSTGRES_PASSWORD: POSTGRES_PASSWORD,
		SERVICE_TABLE:     SERVICE_TABLE,
		REVIEWS_TABLE:     REVIEWS_TABLE,
		REQUEST_TABLE:     REQUEST_TABLE,
		DEBUG:             DEBUG,
		TEST:              TEST,
	}

	return &config, nil
}
