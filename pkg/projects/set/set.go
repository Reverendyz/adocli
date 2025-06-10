package set

import (
	"encoding/json"
	"os"

	"github.com/reverendyz/adocli/pkg/internal/common"
)

func SetProjectInConfigFile(projectId string) error {
	configPath, err := common.GetOrCreateConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var config common.Config
	if len(data) > 0 {
		if err := json.Unmarshal(data, &config); err != nil {
			return err
		}
	}

	config.ProjectId = projectId

	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, updatedData, 0644)
}
