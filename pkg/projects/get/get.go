package get

import (
	"context"
	"fmt"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/utils"
)

func GetProjects(organizationUrl string) {
	coreClient, err := utils.GetClient(organizationUrl)
	if err != nil {
		log.Fatal(err)
	}

	responseValue, err := coreClient.GetProjects(context.Background(), core.GetProjectsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	for responseValue != nil {

		for _, teamProjectReference := range (*responseValue).Value {
			log.Printf("Name[%v] = %v; ID: %v", index, *teamProjectReference.Name, teamProjectReference.Id)
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
			responseValue, err = coreClient.GetProjects(context.Background(), projectArgs)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			responseValue = nil
		}
	}

}
