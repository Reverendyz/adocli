package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/utils"
)

func TestFunc(organizationUrl string) {
	fmt.Println("TestCalled")
	personalAccessToken := utils.GetEnvOrDefault("ADOCLI_PAT", "")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {

		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			index++
		}

		if responseValue.ContinuationToken != "" {

			continuationToken := responseValue.ContinuationToken
			if err != nil {
				log.Fatal(err)
			}

			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			fmt.Println(continuationToken)
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}

}
