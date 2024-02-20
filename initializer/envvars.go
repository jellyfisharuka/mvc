package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvariables() {
	err := godotenv.Load("mvc.env")
	if err!=nil {
		log.Fatal("error loading .env file")
	}
}