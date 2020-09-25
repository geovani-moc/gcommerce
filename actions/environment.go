package actions

import (
	"github.com/joho/godotenv"
)

// LoadEnvVariables in memory
func LoadEnvVariables() (err error) {
	err = godotenv.Load(".env")
	return
}
