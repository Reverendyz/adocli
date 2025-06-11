package create

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
)

func CreateTeam(webApiTeam *core.WebApiTeam) error {
	organizationUrl, err := config.GetFromConfig("organizationUrl")
	if err != nil {
		return err
	}
	projectId, err := config.GetFromConfig("projectId")
	if err != nil {
		return err
	}
	fmt.Println(organizationUrl)
	fmt.Println(projectId)
	coreClient, err := common.GetCoreClient(organizationUrl)
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
