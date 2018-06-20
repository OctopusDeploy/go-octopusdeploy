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

	p := &octopusdeploy.Project{}
	p.LifecycleID = "Lifecycles-1"
	p.Name = "Test Project GoLang"
	p.ProjectGroupID = "ProjectGroups-1"

	createdProject, err := client.Projects.Add(p)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(createdProject.ID)
}
