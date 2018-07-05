package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
)

var octopusURL = os.Getenv("OCTOPUS_URL")
var octopusAPIKey = os.Getenv("OCTOPUS_APIKEY")

func main() {

	httpClient := http.Client{}
	client := octopusdeploy.NewClient(&httpClient, octopusURL, octopusAPIKey)

	p := octopusdeploy.NewProject("Test Project GoLang2", "Lifecycles-1", "ProjectGroups-1")
	err := p.Validate()

	if err != nil {
		fmt.Println(err.Error())
	}

	createdProject, err := client.Project.Add(p)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Created Project ID %s", createdProject.ID)

	project, err := client.Project.Get(createdProject.ID)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(project.Name)
}
