package create

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/logger"
	"go.uber.org/zap"
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

	logger.Info("Team created successfully",
		zap.String("ID", webApiTeamResponse.Id.String()),
		zap.String("Name", *webApiTeamResponse.Name),
	)

	return nil
}
