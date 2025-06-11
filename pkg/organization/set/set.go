package set

import (
	"encoding/json"
	"os"

	"github.com/reverendyz/adocli/config"
)

func SetOrganizationInConfigFile(organizationUrl string) error {
	configPath, err := config.GetOrCreateConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config config.Config
	if len(data) > 0 {
		if err := json.Unmarshal(data, &config); err != nil {
			return err
		}
	}
	config.OrganizationUrl = organizationUrl

	updatedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	os.WriteFile(configPath, updatedData, 0644)

	return nil
}
