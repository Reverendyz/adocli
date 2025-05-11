package get

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/utils"
)

func GetTeams(organizationUrl string, projectId string) {
	coreClient, err := utils.GetClient(organizationUrl)
	if err != nil {
		log.Fatal(err)
	}
	webApiTeamList, err := coreClient.GetTeams(context.Background(), core.GetTeamsArgs{
		ProjectId: &projectId,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, team := range *webApiTeamList {
		log.Printf("Team = %v", team)
	}
}
