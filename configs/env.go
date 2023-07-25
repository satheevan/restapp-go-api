package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	MONGOURI     = "MONGOURI"
	FRONTEND_URI = "FRONTEND_URI"
)

func EnvMongoURI() string {
	return getEnv(MONGOURI)
}

func EnvFrontEndURI() string {
	return getEnv(FRONTEND_URI)
}

func getEnv(str string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading on .env file")
	}
	envData, ok := os.LookupEnv(str)
	if !ok {
		log.Fatalf("Error is getting the value of %s from env file", str)
	}
	return envData
}
