package delete

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/utils"
)

func DeleteTeam(teamId string) error {
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

	err = coreClient.DeleteTeam(context.Background(), core.DeleteTeamArgs{
		TeamId:    &teamId,
		ProjectId: &projectId,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Team %s deleted \n", teamId)

	return nil
}
