package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	fmt.Printf("Loading Environment Variables...")
	err := godotenv.Load()

	if err != nil {
		return err
	}

	fmt.Printf("Environment Variables Loaded !")
	return nil
}
