package get

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/logger"
)

func GetWorkItemById(id string) error {
	client, err := common.GetWIClient(config.OrganizationUrl)
	if err != nil {
		return err
	}

	config.ProjectName, err = common.GetProjectName()
	if err != nil {
		return err
	}

	convertedInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid work item ID '%s': must be a number", id)
	}

	result, err := client.GetWorkItem(context.Background(), workitemtracking.GetWorkItemArgs{
		Id: &convertedInt,
	})
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	logger.Info(string(jsonBytes))

	return nil
}
