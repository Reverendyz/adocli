package utils

import (
	"context"
	"time"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

func GetClient(organizationUrl string) (core.Client, error) {
	personalAccessToken := GetEnvOrDefault("ADOCLI_PAT", "")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return core.NewClient(ctx, connection)
}
