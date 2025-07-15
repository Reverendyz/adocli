package list

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/logger"
	"go.uber.org/zap"
)

func ListAllWorkItemTracking() error {
	client, err := common.GetWIClient(config.OrganizationUrl)
	if err != nil {
		return err
	}
	config.ProjectName, err = common.GetProjectName()
	if err != nil {
		return err
	}
	wiqlQuery := fmt.Sprintf(
		"SELECT [System.Id] FROM WorkItems WHERE [System.TeamProject] = '%s'",
		config.ProjectName,
	)
	result, err := client.QueryByWiql(context.Background(), workitemtracking.QueryByWiqlArgs{
		Wiql: &workitemtracking.Wiql{Query: &wiqlQuery},
	})
	if err != nil {
		return err
	}
	if result.WorkItems == nil || len(*result.WorkItems) == 0 {
		logger.Info("No work items found in this project")
		return nil
	}

	ids := make([]int, 0, len(*result.WorkItems))
	for _, wi := range *result.WorkItems {
		ids = append(ids, *wi.Id)
	}

	batches := splitIDs(ids, 200)
	var allItems []workitemtracking.WorkItem

	for _, batch := range batches {
		items, err := client.GetWorkItems(context.Background(), workitemtracking.GetWorkItemsArgs{
			Ids: &batch,
		})
		if err != nil {
			return err
		}
		allItems = append(allItems, *items...)
	}

	for _, wi := range allItems {
		logger.Info("Work Item",
			zap.Int("ID", *wi.Id),
			zap.String("Title", (*wi.Fields)["System.Title"].(string)),
		)
	}

	return nil
}

func splitIDs(ids []int, batchSize int) [][]int {
	var batches [][]int
	for i := 0; i < len(ids); i += batchSize {
		end := i + batchSize
		if end > len(ids) {
			end = len(ids)
		}
		batches = append(batches, ids[i:end])
	}
	return batches
}
