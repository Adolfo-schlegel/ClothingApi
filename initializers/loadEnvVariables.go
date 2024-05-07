package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	//production environment
	//err := godotenv.Load("/home/dolphin/projects/go/ClothingApi/.env")

	//production environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment file")
	}
}
