package create

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/utils"
)

func CreateTeam(webApiTeam *core.WebApiTeam) error {
	organizationUrl, err := utils.GetFromConfig("organizationUrl")
	if err != nil {
		return err
	}
	projectId, err := utils.GetFromConfig("projectId")
	if err != nil {
		return err
	}
	fmt.Println(organizationUrl)
	fmt.Println(projectId)
	coreClient, err := utils.GetClient(organizationUrl)
	if err != nil {
		return err
	}

	webApiTeamResponse, err := coreClient.CreateTeam(context.Background(), core.CreateTeamArgs{
		Team:      webApiTeam,
		ProjectId: &projectId,
	})
	if err != nil {
		return err
	}

	fmt.Println(*webApiTeamResponse)

	return nil
}
