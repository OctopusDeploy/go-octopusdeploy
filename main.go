package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	//"./client"
	//./model"
)

var serviceUrl = os.Getenv("OCTOPUS_URL")
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

	jsonData, err := model.PrettyJSON(resource)
	fmt.Println(string(jsonData))

	fmt.Println()
}

// CreateSpace creates a test space and outputs the results to the console.
func CreateSpace(client *client.Client) (*model.Space, error) {
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

	space = model.NewSpace(testSpaceName)
	space.SpaceManagersTeams = append(space.SpaceManagersTeams, "teams-administrators")
	space, err = client.Spaces.Add(space)

	if err != nil {
		fmt.Println(err.Error())
		space, err = client.Spaces.GetByName(space.Name)
	}

	jsonData, err := model.PrettyJSON(space)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(jsonData))
	fmt.Println()

	return space, err
}

func CreateProject(client *client.Client) *model.Project {
	fmt.Println("Creating a new project...")

	if client.Projects == nil {
		fmt.Println(fmt.Errorf("unexpected state of client.Projects (nil)"))
		os.Exit(3)
	}

	project := model.NewProject(projectName, lifecycleID, projectGroupID)
	project, err := client.Projects.Add(project)

	if err != nil {
		fmt.Println(err.Error())
		project, err = client.Projects.GetByName(projectName)

		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		jsonData, _ := model.PrettyJSON(project)
		fmt.Println(string(jsonData))
	}

	return project
}

func UpdateProject(client *client.Client, project *model.Project) *model.Project {
	fmt.Println("Updating a project...")

	if client == nil {
		fmt.Println(fmt.Errorf("unexpected state of client (nil)"))
		return nil
	}

	project.Description = "This is the new description..."
	project, err := client.Projects.Update(project)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		jsonData, _ := model.PrettyJSON(project)
		fmt.Println(string(jsonData))
	}

	return project
}

func DeleteProject(client *client.Client, project *model.Project) {
	fmt.Println("Deleting a project...")

	err := client.Projects.Delete(project.ID)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func main() {

	Initialize()

	client, err := client.NewClient(&httpClient, serviceUrl, apiKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	user := model.NewUser("askdhj", "aklsjd")
	user.Password = "asdaasdkhwjerlkqjh987123"
	newUser, err := client.Users.Add(user)
	jsonData, _ := model.PrettyJSON(newUser)
	fmt.Println(string(jsonData))

	authentication, err := client.Authentication.Get()
	jsonData, _ = model.PrettyJSON(authentication)
	fmt.Println(string(jsonData))

	project := CreateProject(client)
	project = UpdateProject(client, project)
	DeleteProject(client, project)

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

	err = client.Spaces.Delete(updatedSpace.ID)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Space deleted.")

	p := model.NewProject("Test Project GoLang2", "Lifecycles-1", "ProjectGroups-1")

	if client.Projects == nil {
		fmt.Println(fmt.Errorf("unexpected state of client.Projects (nil)"))
		os.Exit(3)
	}

	project, err = client.Projects.Add(p)

	if err != nil {
		fmt.Println(err.Error())
	} else { //This isn't idomatic go, but it allows the demo to continue if the create fails
		fmt.Printf("Created Project ID %s", project.ID)

		project, err := client.Projects.Get(project.ID)

		if err != nil {
			fmt.Println(err.Error())
		} else { //This isn't idomatic go, but it allows the demo to continue if the create fails
			fmt.Println(project.Name)
		}
	}

	e := model.NewEnvironment("Test Environment GoLang", "Test environment created by go-octopusdeploy", false)
	err = e.Validate()

	if err != nil {
		fmt.Println(err.Error())
	}

	createdEnvironment, err := client.Environments.Add(e)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Created Project ID %s", createdEnvironment.ID)
		environment, err := client.Environments.Get(createdEnvironment.ID)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(environment.Name)
	}
}

func Initialize() {
	proxyStr := "http://127.0.0.1:5555"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	httpClient = http.Client{Transport: tr}
}
