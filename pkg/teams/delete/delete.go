package delete

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/pkg/internal/common"
)

func DeleteTeam(teamId string) error {
	organizationUrl, err := common.GetFromConfig("organizationUrl")
	if err != nil {
		return err
	}
	projectId, err := common.GetFromConfig("projectId")
	if err != nil {
		return err
	}
	coreClient, err := common.GetClient(organizationUrl)
	if err != nil {
		return err
	}

	team, err := coreClient.GetTeam(context.Background(), core.GetTeamArgs{
		ProjectId: &projectId,
		TeamId:    &teamId,
	})
	if err != nil {
		return err
	}

	teamId = team.Id.String()

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
