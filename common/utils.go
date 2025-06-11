package common

import (
	"os"
	"strings"

	"github.com/reverendyz/adocli/config"
	"github.com/spf13/cobra"
)

func GetEnvOrDefault(envVarName string, fallbackValue string) string {
	envVarContent := os.Getenv(strings.ToUpper(envVarName))
	if envVarContent == "" {
		return fallbackValue
	}

	return envVarContent
}

func SetOrLoadFlag(cmd *cobra.Command, flagName string, fallbackKey string) (string, error) {
	val, err := cmd.Flags().GetString(flagName)

	if err == nil && val != "" {
		return val, nil
	}
	return config.GetFromConfig(fallbackKey)
}
