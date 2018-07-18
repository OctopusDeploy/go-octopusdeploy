package octopusdeploy

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Octopus Deploy API Services
	DeploymentProcess *DeploymentProcessService
	ProjectGroup      *ProjectGroupService
	Project           *ProjectService
	ProjectTrigger    *ProjectTriggerService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, octopusURL, octopusAPIKey string) *Client {
	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api/", baseURLWithAPI)
	fmt.Println(baseURLWithAPI)
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", octopusAPIKey)
	return &Client{
		sling:             base,
		DeploymentProcess: NewDeploymentProcessService(base.New()),
		ProjectGroup:      NewProjectGroupService(base.New()),
		Project:           NewProjectService(base.New()),
		ProjectTrigger:    NewProjectTriggerService(base.New()),
	}
}

type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}

func APIErrorChecker(urlPath string, resp *http.Response, wantedResponseCode int, slingError error, octopusDeployError *APIError) error {
	if octopusDeployError.Errors != nil {
		return fmt.Errorf("cannot get all projects. response from octopusdeploy %s: ", octopusDeployError.Errors)
	}

	if slingError != nil {
		return fmt.Errorf("cannot get path %s from server. failure from http client %v", urlPath, slingError)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrItemNotFound
	}

	if resp.StatusCode != wantedResponseCode {
		return fmt.Errorf("cannot get item from %s from server. response from server %s", urlPath, resp.Status)
	}

	return nil
}

// Generic OctopusDeploy API Get Function
func apiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Get(path).Receive(inputStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

// Generic OctopusDeploy API Add Function
func apiAdd(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Post(path).BodyJSON(inputStruct).Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Add Function
func apiUpdate(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Put(path).BodyJSON(inputStruct).Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Delete Function
func apiDelete(sling *sling.Sling, path string) (error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Delete(path).Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

var ErrItemNotFound = errors.New("cannot find the item")
