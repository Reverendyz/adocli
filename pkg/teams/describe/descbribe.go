package describe

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
		log.Println("Team Details:")
		log.Printf("  ID: %s", team.Id.String())

		if team.Name != nil {
			log.Printf("  Name: %s", *team.Name)
		} else {
			log.Println("  Name: <nil>")
		}

		if team.Url != nil {
			log.Printf("  URL: %s", *team.Url)
		} else {
			log.Println("  URL: <nil>")
		}

		if team.Description != nil {
			log.Printf("  Description: %s", *team.Description)
		} else {
			log.Println("  Description: <nil>")
		}

		if team.IdentityUrl != nil {
			log.Printf("  Identity URL: %s", *team.IdentityUrl)
		} else {
			log.Println("  Identity URL: <nil>")
		}

		if team.ProjectName != nil {
			log.Printf("  Project Name: %s", *team.ProjectName)
		} else {
			log.Println("  Project Name: <nil>")
		}

		log.Printf("  Project ID: %s", team.ProjectId.String())

		if team.Url != nil {
			log.Printf("  Project URL: %s", *team.Url)
		} else {
			log.Println("  Project URL: <nil>")
		}
	}
}
