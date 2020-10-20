package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

var octopusURL = os.Getenv("OCTOPUS_URL")
var apiKey = os.Getenv("OCTOPUS_APIKEY")
var testSpaceName = "Test Space"
var projectName = "Test Project"
var lifecycleID = "Lifecycles-1"
var projectGroupID = "ProjectGroups-1"
var httpClient http.Client

// OutputAsJSON outputs a resource to the console as JSON.
func OutputAsJSON(resource interface{}, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonData := octopusdeploy.PrettyJSON(resource)
	fmt.Println(jsonData)
	fmt.Println()
}

// CreateSpace creates a test space and outputs the results to the console.
func CreateSpace(client *octopusdeploy.Client) (*octopusdeploy.Space, error) {
	fmt.Println("Creating a new space...")

	if client.Spaces == nil {
		fmt.Println(fmt.Errorf("unexpected state of client.Spaces (nil)"))
		os.Exit(3)
	}

	space, err := client.Spaces.GetByName(testSpaceName)

	if err == nil {
		fmt.Println("Space already exists.")
		return space, err
	}

	space = octopusdeploy.NewSpace(testSpaceName)
	space.SpaceManagersTeams = append(space.SpaceManagersTeams, "teams-administrators")
	space, err = client.Spaces.Add(space)

	if err != nil {
		fmt.Println(err.Error())
		space, err = client.Spaces.GetByName(space.Name)

		if err != nil {
			return nil, err

		}
	}

	jsonData := octopusdeploy.PrettyJSON(space)
	fmt.Println(jsonData)
	fmt.Println()

	return space, err
}

func createProject(client *octopusdeploy.Client) *octopusdeploy.Project {
	fmt.Println("Creating a new project...")

	if client.Projects == nil {
		fmt.Println(fmt.Errorf("unexpected state of client.Projects (nil)"))
		os.Exit(3)
	}

	project := octopusdeploy.NewProject(projectName, lifecycleID, projectGroupID)
	project, err := client.Projects.Add(project)

	if err != nil {
		fmt.Println(err.Error())
		project, err = client.Projects.GetByName(projectName)

		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		jsonData := octopusdeploy.PrettyJSON(project)
		fmt.Println(jsonData)
	}

	return project
}

func updateProject(client *octopusdeploy.Client, project *octopusdeploy.Project) *octopusdeploy.Project {
	fmt.Println("Updating a project...")

	if client == nil {
		fmt.Println(fmt.Errorf("unexpected state of client (nil)"))
		return nil
	}

	project.Description = "This is the new description..."
	project, err := client.Projects.Update(*project)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		jsonData := octopusdeploy.PrettyJSON(project)
		fmt.Println(jsonData)
	}

	return project
}

func deleteProject(client *octopusdeploy.Client, project *octopusdeploy.Project) {
	fmt.Println("Deleting a project...")

	err := client.Projects.DeleteByID(project.ID)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func main() {

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	user := octopusdeploy.NewUser("askdhj", "aklsjd")
	user.Password = "asdaasdkhwjerlkqjh987123"

	newUser, err := client.Users.Add(user)
	if err != nil {
		fmt.Println(err)
	}

	jsonData := octopusdeploy.PrettyJSON(newUser)
	fmt.Println(jsonData)

	authentication, err := client.Authentication.Get()
	if err != nil {
		fmt.Println(err)
	}

	jsonData = octopusdeploy.PrettyJSON(authentication)
	fmt.Println(jsonData)

	project := createProject(client)
	project = updateProject(client, project)
	deleteProject(client, project)

	OutputAsJSON(client.Accounts.GetAll())
	OutputAsJSON(client.ActionTemplates.GetAll())
	OutputAsJSON(client.Certificates.GetAll())
	OutputAsJSON(client.Channels.GetAll())
	OutputAsJSON(client.DeploymentProcesses.GetAll())
	OutputAsJSON(client.Environments.GetAll())
	OutputAsJSON(client.Feeds.GetAll())
	OutputAsJSON(client.Interruptions.GetAll())
	OutputAsJSON(client.LibraryVariableSets.GetAll())
	OutputAsJSON(client.Lifecycles.GetAll())
	OutputAsJSON(client.Machines.GetAll())
	OutputAsJSON(client.MachinePolicies.GetAll())
	OutputAsJSON(client.Projects.GetAll())
	OutputAsJSON(client.ProjectGroups.GetAll())
	OutputAsJSON(client.ProjectTriggers.GetAll())
	OutputAsJSON(client.Spaces.GetAll())
	OutputAsJSON(client.TagSets.GetAll())
	OutputAsJSON(client.Tenants.GetAll())
	OutputAsJSON(client.Users.GetAll())

	space, err := CreateSpace(client)

	if err != nil {
		fmt.Println(err.Error())
	}

	// delete a space
	fmt.Println("Deleting an existing space...")

	// stop the task queue before deleting the space
	fmt.Print("Stopping task queue... ")
	space.TaskQueueStopped = true
	updatedSpace, err := client.Spaces.Update(space)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("done.")
	}

	err = client.Spaces.DeleteByID(updatedSpace.ID)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Space deleted.")

	p := octopusdeploy.NewProject("Test Project GoLang2", "Lifecycles-1", "ProjectGroups-1")

	if client.Projects == nil {
		fmt.Println(fmt.Errorf("unexpected state of client.Projects (nil)"))
		os.Exit(3)
	}

	project, err = client.Projects.Add(p)

	if err != nil {
		fmt.Println(err.Error())
	} else { //This isn't idomatic go, but it allows the demo to continue if the create fails
		fmt.Printf("Created Project ID %s", project.ID)

		project, err := client.Projects.GetByID(project.ID)

		if err != nil {
			fmt.Println(err.Error())
		} else { //This isn't idomatic go, but it allows the demo to continue if the create fails
			fmt.Println(project.Name)
		}
	}

	e := octopusdeploy.NewEnvironment("Test Environment (OK to Delete)")
	e.Description = "Test environment created by go-octopusdeploy"
	e.UseGuidedFailure = false
	err = e.Validate()

	if err != nil {
		fmt.Println(err.Error())
	}

	createdEnvironment, err := client.Environments.Add(e)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Created Project ID %s", createdEnvironment.ID)
		environment, err := client.Environments.GetByID(createdEnvironment.ID)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(environment.Name)
	}
}
