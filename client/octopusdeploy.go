package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// Client is an OctopusDeploy for making OctpusDeploy API requests.
type Client struct {
	sling               *sling.Sling
	Accounts            *AccountService
	ActionTemplates     *ActionTemplateService
	APIKeys             *APIKeyService
	Authentication      *AuthenticationService
	Certificates        *CertificateService
	Channels            *ChannelService
	Configuration       *ConfigurationService
	DeploymentProcesses *DeploymentProcessService
	Environments        *EnvironmentService
	Feeds               *FeedService
	Interruptions       *InterruptionsService
	LibraryVariableSets *LibraryVariableSetService
	Lifecycles          *LifecycleService
	Machines            *MachineService
	MachinePolicies     *MachinePolicyService
	Projects            *ProjectService
	ProjectGroups       *ProjectGroupService
	ProjectTriggers     *ProjectTriggerService
	Root                *RootService
	Spaces              *SpaceService
	TagSets             *TagSetService
	Tenants             *TenantService
	Users               *UserService
	Variables           *VariableService
}

// NewClient returns a new
func NewClient(httpClient *http.Client, octopusURL string, apiKey string, spaceName string) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	if isEmpty(octopusURL) {
		return nil, createInvalidParameterError("NewClient", "octopusURL")
	}

	if isEmpty(apiKey) {
		return nil, createInvalidParameterError("NewClient", "apiKey")
	}

	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api", baseURLWithAPI)

	// fetch root resource and process paths
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", apiKey)
	rootService := NewRootService(base.New(), baseURLWithAPI)
	root, err := rootService.Get()

	if err != nil {
		return nil, err
	}

	accountsPath := root.Links["Accounts"]
	actionTemplatesPath := root.Links["ActionTemplates"]
	apiKeysPath := "/api/users"
	authenticationPath := root.Links["Authentication"]
	certificatesPath := root.Links["Certificates"]
	channelsPath := root.Links["Channels"]
	configurationPath := root.Links["Configuration"]
	deploymentProcessesPath := root.Links["DeploymentProcesses"]
	environmentsPath := root.Links["Environments"]
	feedsPath := root.Links["Feeds"]
	interruptionsPath := root.Links["Interruptions"]
	machinesPath := root.Links["Machines"]
	machinePoliciesPath := root.Links["MachinePolicies"]
	libraryVariableSetsPath := root.Links["LibraryVariableSets"]
	lifecyclesPath := root.Links["Lifecycles"]
	projectsPath := root.Links["Projects"]
	projectGroupsPath := root.Links["ProjectGroups"]
	projectTriggersPath := root.Links["ProjectTriggers"]
	spacesPath := root.Links["Spaces"]
	tagSetsPath := root.Links["TagSets"]
	tenantsPath := root.Links["Tenants"]
	usersPath := root.Links["Users"]
	variablesPath := root.Links["Variables"]

	if !isEmpty(spaceName) {
		baseURLWithAPI = fmt.Sprintf("%s/%s", baseURLWithAPI, spaceName)
		base = sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", apiKey)
		rootService = NewRootService(base.New(), baseURLWithAPI)
		root, err = rootService.Get()

		if err != nil {
			return nil, err
		}

		if !isEmpty(root.Links["Accounts"]) {
			accountsPath = root.Links["Accounts"]
		}

		if !isEmpty(root.Links["ActionTemplates"]) {
			actionTemplatesPath = root.Links["ActionTemplates"]
		}

		if !isEmpty(root.Links["Authentication"]) {
			authenticationPath = root.Links["Authentication"]
		}

		if !isEmpty(root.Links["Authentication"]) {
			authenticationPath = root.Links["Authentication"]
		}

		if !isEmpty(root.Links["Certificates"]) {
			certificatesPath = root.Links["Certificates"]
		}

		if !isEmpty(root.Links["Channels"]) {
			channelsPath = root.Links["Channels"]
		}

		if !isEmpty(root.Links["Configuration"]) {
			configurationPath = root.Links["Configuration"]
		}

		if !isEmpty(root.Links["DeploymentProcesses"]) {
			deploymentProcessesPath = root.Links["DeploymentProcesses"]
		}

		if !isEmpty(root.Links["Environments"]) {
			environmentsPath = root.Links["Environments"]
		}

		if !isEmpty(root.Links["Feeds"]) {
			feedsPath = root.Links["Feeds"]
		}

		if !isEmpty(root.Links["Interruptions"]) {
			interruptionsPath = root.Links["Interruptions"]
		}

		if !isEmpty(root.Links["Machines"]) {
			machinesPath = root.Links["Machines"]
		}

		if !isEmpty(root.Links["MachinePolicies"]) {
			machinePoliciesPath = root.Links["MachinePolicies"]
		}

		if !isEmpty(root.Links["LibraryVariableSets"]) {
			libraryVariableSetsPath = root.Links["LibraryVariableSets"]
		}

		if !isEmpty(root.Links["Lifecycles"]) {
			lifecyclesPath = root.Links["Lifecycles"]
		}

		if !isEmpty(root.Links["Projects"]) {
			projectsPath = root.Links["Projects"]
		}

		if !isEmpty(root.Links["ProjectGroups"]) {
			projectGroupsPath = root.Links["ProjectGroups"]
		}

		if !isEmpty(root.Links["ProjectTriggers"]) {
			projectTriggersPath = root.Links["ProjectTriggers"]
		}

		if !isEmpty(root.Links["Spaces"]) {
			spacesPath = root.Links["Spaces"]
		}

		if !isEmpty(root.Links["TagSets"]) {
			tagSetsPath = root.Links["TagSets"]
		}

		if !isEmpty(root.Links["Tenants"]) {
			tenantsPath = root.Links["Tenants"]
		}

		if !isEmpty(root.Links["Users"]) {
			usersPath = root.Links["Users"]
		}

		if !isEmpty(root.Links["Variables"]) {
			variablesPath = root.Links["Variables"]
		}
	}

	return &Client{
		sling:               base,
		Accounts:            NewAccountService(base.New(), accountsPath),
		ActionTemplates:     NewActionTemplateService(base.New(), actionTemplatesPath),
		APIKeys:             NewAPIKeyService(base.New(), apiKeysPath),
		Authentication:      NewAuthenticationService(base.New(), authenticationPath),
		Certificates:        NewCertificateService(base.New(), certificatesPath),
		Channels:            NewChannelService(base.New(), channelsPath),
		Configuration:       NewConfigurationService(base.New(), configurationPath),
		DeploymentProcesses: NewDeploymentProcessService(base.New(), deploymentProcessesPath),
		Environments:        NewEnvironmentService(base.New(), environmentsPath),
		Feeds:               NewFeedService(base.New(), feedsPath),
		Interruptions:       NewInterruptionsService(base.New(), interruptionsPath),
		Machines:            NewMachineService(base.New(), machinesPath),
		MachinePolicies:     NewMachinePolicyService(base.New(), machinePoliciesPath),
		LibraryVariableSets: NewLibraryVariableSetService(base.New(), libraryVariableSetsPath),
		Lifecycles:          NewLifecycleService(base.New(), lifecyclesPath),
		Projects:            NewProjectService(base.New(), projectsPath),
		ProjectGroups:       NewProjectGroupService(base.New(), projectGroupsPath),
		ProjectTriggers:     NewProjectTriggerService(base.New(), projectTriggersPath),
		Root:                rootService,
		Spaces:              NewSpaceService(base.New(), spacesPath),
		TagSets:             NewTagSetService(base.New(), tagSetsPath),
		Tenants:             NewTenantService(base.New(), tenantsPath),
		Users:               NewUserService(base.New(), usersPath),
		Variables:           NewVariableService(base.New(), variablesPath),
	}, nil
}

type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}

// APIErrorChecker is a generic error handler for the OctopusDeploy API.
func APIErrorChecker(urlPath string, resp *http.Response, wantedResponseCode int, slingError error, octopusDeployError *APIError) error {
	if octopusDeployError.Errors != nil {
		return fmt.Errorf("octopus deploy api returned an error on endpoint %s - %s", urlPath, octopusDeployError.Errors)
	}

	if slingError != nil {
		return fmt.Errorf("cannot get endpoint %s from server. failure from http client %v", urlPath, slingError)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrItemNotFound
	}

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	if resp.StatusCode != wantedResponseCode {
		return fmt.Errorf("cannot get item from endpoint %s. response from server %s", urlPath, resp.Status)
	}

	return nil
}

// LoadNextPage checks if the next page should be loaded from the API. Returns the new path and a bool if the next page should be checked.
func LoadNextPage(pagedResults model.PagedResults) (string, bool) {
	if pagedResults.Links.PageNext != "" {
		return pagedResults.Links.PageNext, true
	}

	return "", false
}

// Generic OctopusDeploy API Get Function.
func apiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError("apiGet", "sling")
	}

	getClient := sling.New()

	if getClient == nil {
		return nil, createClientInitializationError("apiGet")
	}

	getClient = getClient.Get(path)

	if getClient == nil {
		return nil, createClientInitializationError("apiGet")
	}

	octopusDeployError := new(APIError)
	resp, err := getClient.Receive(inputStruct, &octopusDeployError)

	if err != nil {
		return nil, err
	}

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func apiAdd(sling *sling.Sling, inputStruct, resource model.ResourceInterface, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError("apiAdd", "sling")
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError("apiAdd", "path")
	}

	postClient := sling.New()

	if postClient == nil {
		return nil, createClientInitializationError("apiAdd")
	}

	postClient = postClient.Post(path)

	if postClient == nil {
		return nil, createClientInitializationError("apiAdd")
	}

	request := postClient.BodyJSON(inputStruct)

	if request == nil {
		return nil, createClientInitializationError("apiAdd")
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// apiPost post to octopus and expect a 200 response code.
func apiPost(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError("apiPost", "sling")
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError("apiPost", "path")
	}

	postClient := sling.New()

	if postClient == nil {
		return nil, createClientInitializationError("apiPost")
	}

	postClient = postClient.Post(path)

	if postClient == nil {
		return nil, createClientInitializationError("apiPost")
	}

	request := postClient.BodyJSON(inputStruct)

	if request == nil {
		return nil, createClientInitializationError("apiPost")
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Update Function.
func apiUpdate(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError("apiUpdate", "sling")
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError("apiUpdate", "path")
	}

	putClient := sling.New()

	if putClient == nil {
		return nil, createClientInitializationError("apiUpdate")
	}

	putClient = putClient.Put(path)

	if putClient == nil {
		return nil, createClientInitializationError("apiUpdate")
	}

	request := putClient.BodyJSON(inputStruct)

	if request == nil {
		return nil, createClientInitializationError("apiUpdate")
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Delete Function.
func apiDelete(sling *sling.Sling, path string) error {
	if sling == nil {
		return createInvalidParameterError("apiDelete", "sling")
	}

	if isEmpty(path) {
		return createInvalidParameterError("apiDelete", "path")
	}

	deleteClient := sling.New()

	if deleteClient == nil {
		return createClientInitializationError("apiDelete")
	}

	deleteClient = deleteClient.Delete(path)

	if deleteClient == nil {
		return createClientInitializationError("apiDelete")
	}

	octopusDeployError := new(APIError)
	resp, err := deleteClient.Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")
