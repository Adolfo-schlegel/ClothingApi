package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	//err := godotenv.Load("/home/dolphin/projects/go/ClothingApi/.env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment file")
	}
}
