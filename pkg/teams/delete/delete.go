package delete

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/logger"
	"go.uber.org/zap"
)

func DeleteTeam(teamId string) error {
	organizationUrl, err := config.GetFromConfig("organizationUrl")
	if err != nil {
		return err
	}
	projectId, err := config.GetFromConfig("projectId")
	if err != nil {
		return err
	}
	coreClient, err := common.GetCoreClient(organizationUrl)
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

	logger.Info("Team deleted successfully", zap.String("TeamID", teamId))

	return nil
}
