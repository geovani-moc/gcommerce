package util

import (
	"errors"

	"github.com/joho/godotenv"
)

// LoadEnvVariables in memory
func LoadEnvVariables() error {

	err := godotenv.Load(".env")
	if nil != err {
		return errors.New("env: nao foi localizado nehum arquivo com variaveis de ambiente")
	}

	return nil
}
