package initializers

import (
	"log"

	broEnv "github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := broEnv.Load()

	if err != nil {
		log.Fatal("Bro Something Is Wrong 😑")
	}
}
