package list

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/logger"
	"go.uber.org/zap"
)

func ProjectsList(organizationUrl string) error {
	coreClient, err := common.GetCoreClient(organizationUrl)
	if err != nil {
		return err
	}

	responseValue, err := coreClient.GetProjects(context.Background(), core.GetProjectsArgs{})
	if err != nil {
		return err
	}

	index := 0
	for responseValue != nil {

		for _, teamProjectReference := range (*responseValue).Value {
			logger.Info("Project",
				zap.Int("Index", index),
				zap.String("Name", *teamProjectReference.Name),
				zap.String("ID", teamProjectReference.Id.String()),
			)
			index++
		}

		if responseValue.ContinuationToken != "" {
			continuationToken := responseValue.ContinuationToken

			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			responseValue, err = coreClient.GetProjects(context.Background(), projectArgs)
			if err != nil {
				return err
			}
		} else {
			responseValue = nil
		}
	}

	return nil
}
