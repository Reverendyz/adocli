package common

import (
	"context"
	"time"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/reverendyz/adocli/config"
)

func GetCoreClient(organizationUrl string) (core.Client, error) {
	personalAccessToken := GetEnvOrDefault("ADOCLI_PAT", "")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return core.NewClient(ctx, connection)
}

func GetWIClient(organizationUrl string) (workitemtracking.Client, error) {
	personalAccessToken := GetEnvOrDefault("ADOCLI_PAT", "")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return workitemtracking.NewClient(ctx, connection)
}

func GetProjectName() (string, error) {
	client, err := GetCoreClient(config.OrganizationUrl)
	if err != nil {
		return "", err
	}
	teamProject, err := client.GetProject(context.Background(), core.GetProjectArgs{
		ProjectId: &config.ProjectId,
	})
	if err != nil {
		return "", err
	}
	return *teamProject.Name, nil
}
