package utils

import (
	"os"
	"strings"
)

func GetEnvOrDefault(envVarName string, fallbackValue string) string {
	envVarContent := os.Getenv(strings.ToUpper(envVarName))
	if envVarContent == "" {
		return fallbackValue
	}

	return envVarContent
}
