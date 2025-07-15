package describe

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/logger"
	"go.uber.org/zap"
)

func GetTeams(organizationUrl string, projectId string) error {
	coreClient, err := common.GetCoreClient(organizationUrl)
	if err != nil {
		return err
	}
	webApiTeamList, err := coreClient.GetTeams(context.Background(), core.GetTeamsArgs{
		ProjectId: &projectId,
	})
	if err != nil {
		return err
	}

	for _, team := range *webApiTeamList {
		fields := []zap.Field{
			zap.String("ID", team.Id.String()),
		}

		if team.Name != nil {
			fields = append(fields, zap.String("Name", *team.Name))
		}

		if team.Url != nil {
			fields = append(fields, zap.String("URL", *team.Url))
		}

		if team.Description != nil {
			fields = append(fields, zap.String("Description", *team.Description))
		}

		if team.IdentityUrl != nil {
			fields = append(fields, zap.String("IdentityURL", *team.IdentityUrl))
		}

		if team.ProjectName != nil {
			fields = append(fields, zap.String("ProjectName", *team.ProjectName))
		}

		fields = append(fields, zap.String("ProjectID", team.ProjectId.String()))

		logger.Info("Team Details", fields...)
	}

	return nil
}
