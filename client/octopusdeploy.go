package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// Client is an OctopusDeploy for making Octopus API requests.
type Client struct {
	sling               *sling.Sling
	Accounts            *accountService
	ActionTemplates     *actionTemplateService
	APIKeys             *apiKeyService
	Authentication      *authenticationService
	Certificates        *certificateService
	Channels            *channelService
	Configuration       *configurationService
	DeploymentProcesses *deploymentProcessService
	Deployments         *deploymentService
	Environments        *environmentService
	Feeds               *feedService
	Interruptions       *interruptionsService
	LibraryVariableSets *libraryVariableSetService
	Lifecycles          *lifecycleService
	Machines            *machineService
	MachinePolicies     *machinePolicyService
	Projects            *projectService
	ProjectGroups       *projectGroupService
	ProjectTriggers     *projectTriggerService
	Root                *rootService
	Spaces              *spaceService
	TagSets             *tagSetService
	Tenants             *tenantService
	Users               *userService
	Variables           *variableService
}

// NewClient returns a new Octopus API client. If a nil client is provided, a
// new http.Client will be used.
func NewClient(httpClient *http.Client, octopusURL string, apiKey string, spaceID string) (*Client, error) {
	if isEmpty(octopusURL) {
		return nil, createInvalidParameterError(clientNewClient, parameterOctopusURL)
	}

	if isEmpty(apiKey) {
		return nil, createInvalidParameterError(clientNewClient, parameterAPIKey)
	}

	if !isAPIKey(apiKey) {
		return nil, createInvalidParameterError(clientNewClient, parameterAPIKey)
	}

	_, err := url.Parse(octopusURL)
	if err != nil {
		return nil, createInvalidParameterError(clientNewClient, parameterOctopusURL)
	}

	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api", baseURLWithAPI)

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// fetch root resource and process paths
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set(clientAPIKeyHTTPHeader, apiKey)
	rootService := newRootService(base, baseURLWithAPI)

	root, err := rootService.Get()
	if err != nil {
		return nil, err
	}

	accountsPath := root.Links[linkAccounts]
	actionTemplatesPath := root.Links[linkActionTemplates]
	actionTemplatesCategoriesURL, _ := url.Parse(root.Links[linkActionTemplatesCategories])
	actionTemplatesSearchURL, _ := url.Parse(root.Links[linkActionTemplatesSearch])
	actionTemplateVersionedLogoURL, _ := url.Parse(root.Links[linkActionTemplateVersionedLogo])
	apiKeysPath := "/api/users"
	authenticationPath := root.Links[linkAuthentication]
	certificatesPath := root.Links[linkCertificates]
	channelsPath := root.Links[linkChannels]
	configurationPath := root.Links[linkConfiguration]
	deploymentProcessesPath := root.Links[linkDeploymentProcesses]
	deploymentsPath := root.Links[linkDeployments]
	environmentsPath := root.Links[linkEnvironments]
	feedsPath := root.Links[linkFeeds]
	interruptionsPath := root.Links[linkInterruptions]
	machinesPath := root.Links[linkMachines]
	machinePoliciesPath := root.Links[linkMachinePolicies]
	libraryVariableSetsPath := root.Links[linkLibraryVariableSets]
	lifecyclesPath := root.Links[linkLifecycles]
	projectsPath := root.Links[linkProjects]
	projectGroupsPath := root.Links[linkProjectGroups]
	projectTriggersPath := root.Links[linkProjectTriggers]
	rootPath := root.Links[linkSelf]
	spacesPath := root.Links[linkSpaces]
	tagSetsPath := root.Links[linkTagSets]
	tenantsPath := root.Links[linkTenants]
	usersPath := root.Links[linkUsers]
	variablesPath := root.Links[linkVariables]

	if !isEmpty(spaceID) {
		baseURLWithAPI = fmt.Sprintf("%s/%s", baseURLWithAPI, spaceID)
		base = sling.New().Client(httpClient).Base(baseURLWithAPI).Set(clientAPIKeyHTTPHeader, apiKey)
		rootService = newRootService(base, baseURLWithAPI)
		root, err = rootService.Get()

		if err != nil {
			if err == ErrItemNotFound {
				return nil, fmt.Errorf("the space ID (%s) cannot be found", spaceID)
			}
			return nil, err
		}

		if !isEmpty(root.Links[linkAccounts]) {
			accountsPath = root.Links[linkAccounts]
		}

		if !isEmpty(root.Links[linkActionTemplates]) {
			actionTemplatesPath = root.Links[linkActionTemplates]
		}

		if !isEmpty(root.Links[linkActionTemplatesCategories]) {
			actionTemplatesCategoriesURL, _ = url.Parse(root.Links[linkActionTemplatesCategories])
		}

		if !isEmpty(root.Links[linkActionTemplatesSearch]) {
			actionTemplatesSearchURL, _ = url.Parse(root.Links[linkActionTemplatesSearch])
		}

		if !isEmpty(root.Links[linkActionTemplateVersionedLogo]) {
			actionTemplateVersionedLogoURL, _ = url.Parse(root.Links[linkActionTemplateVersionedLogo])
		}

		if !isEmpty(root.Links[linkAuthentication]) {
			authenticationPath = root.Links[linkAuthentication]
		}

		if !isEmpty(root.Links[linkCertificates]) {
			certificatesPath = root.Links[linkCertificates]
		}

		if !isEmpty(root.Links[linkChannels]) {
			channelsPath = root.Links[linkChannels]
		}

		if !isEmpty(root.Links[linkConfiguration]) {
			configurationPath = root.Links[linkConfiguration]
		}

		if !isEmpty(root.Links[linkDeploymentProcesses]) {
			deploymentProcessesPath = root.Links[linkDeploymentProcesses]
		}

		if !isEmpty(root.Links[linkDeployments]) {
			deploymentsPath = root.Links[linkDeployments]
		}

		if !isEmpty(root.Links[linkEnvironments]) {
			environmentsPath = root.Links[linkEnvironments]
		}

		if !isEmpty(root.Links[linkFeeds]) {
			feedsPath = root.Links[linkFeeds]
		}

		if !isEmpty(root.Links[linkInterruptions]) {
			interruptionsPath = root.Links[linkInterruptions]
		}

		if !isEmpty(root.Links[linkMachines]) {
			machinesPath = root.Links[linkMachines]
		}

		if !isEmpty(root.Links[linkMachinePolicies]) {
			machinePoliciesPath = root.Links[linkMachinePolicies]
		}

		if !isEmpty(root.Links[linkLibraryVariableSets]) {
			libraryVariableSetsPath = root.Links[linkLibraryVariableSets]
		}

		if !isEmpty(root.Links[linkLifecycles]) {
			lifecyclesPath = root.Links[linkLifecycles]
		}

		if !isEmpty(root.Links[linkProjects]) {
			projectsPath = root.Links[linkProjects]
		}

		if !isEmpty(root.Links[linkProjectGroups]) {
			projectGroupsPath = root.Links[linkProjectGroups]
		}

		if !isEmpty(root.Links[linkProjectTriggers]) {
			projectTriggersPath = root.Links[linkProjectTriggers]
		}

		if !isEmpty(root.Links[linkSelf]) {
			rootPath = root.Links[linkSelf]
		}

		if !isEmpty(root.Links[linkSpaces]) {
			spacesPath = root.Links[linkSpaces]
		}

		if !isEmpty(root.Links[linkTagSets]) {
			tagSetsPath = root.Links[linkTagSets]
		}

		if !isEmpty(root.Links[linkTenants]) {
			tenantsPath = root.Links[linkTenants]
		}

		if !isEmpty(root.Links[linkUsers]) {
			usersPath = root.Links[linkUsers]
		}

		if !isEmpty(root.Links[linkVariables]) {
			variablesPath = root.Links[linkVariables]
		}
	}

	return &Client{
		sling:               base,
		Accounts:            newAccountService(base, accountsPath),
		ActionTemplates:     newActionTemplateService(base, actionTemplatesPath, *actionTemplatesCategoriesURL, *actionTemplatesSearchURL, *actionTemplateVersionedLogoURL),
		APIKeys:             newAPIKeyService(base, apiKeysPath),
		Authentication:      newAuthenticationService(base, authenticationPath),
		Certificates:        newCertificateService(base, certificatesPath),
		Channels:            newChannelService(base, channelsPath),
		Configuration:       newConfigurationService(base, configurationPath),
		Deployments:         newDeploymentService(base, deploymentsPath),
		DeploymentProcesses: newDeploymentProcessService(base, deploymentProcessesPath),
		Environments:        newEnvironmentService(base, environmentsPath),
		Feeds:               newFeedService(base, feedsPath),
		Interruptions:       newInterruptionsService(base, interruptionsPath),
		Machines:            newMachineService(base, machinesPath),
		MachinePolicies:     newMachinePolicyService(base, machinePoliciesPath),
		LibraryVariableSets: newLibraryVariableSetService(base, libraryVariableSetsPath),
		Lifecycles:          newLifecycleService(base, lifecyclesPath),
		Projects:            newProjectService(base, projectsPath),
		ProjectGroups:       newProjectGroupService(base, projectGroupsPath),
		ProjectTriggers:     newProjectTriggerService(base, projectTriggersPath),
		Root:                newRootService(base, rootPath),
		Spaces:              newSpaceService(base, spacesPath),
		TagSets:             newTagSetService(base, tagSetsPath),
		Tenants:             newTenantService(base, tenantsPath),
		Users:               newUserService(base, usersPath),
		Variables:           newVariableService(base, variablesPath),
	}, nil
}

// APIError is a generic structure for containing errors for API operations.
type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

// Error creates a predefined error for Octopus API responses.
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
	if pagedResults.Links.PageNext != emptyString {
		return pagedResults.Links.PageNext, true
	}

	return emptyString, false
}

// Generic OctopusDeploy API Get Function.
func apiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(operationAPIGet, parameterSling)
	}

	getClient := sling.Get(path)
	if getClient == nil {
		return nil, createClientInitializationError(operationAPIGet)
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
func apiAdd(sling *sling.Sling, inputStruct, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(operationAPIAdd, parameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(operationAPIAdd, parameterPath)
	}

	postClient := sling.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(operationAPIAdd)
	}

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(operationAPIAdd)
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
		return nil, createInvalidParameterError(operationAPIPost, parameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(operationAPIPost, parameterPath)
	}

	postClient := sling.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(operationAPIPost)
	}

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(operationAPIPost)
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
		return nil, createInvalidParameterError(operationAPIUpdate, parameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(operationAPIUpdate, parameterPath)
	}

	putClient := sling.Put(path)
	if putClient == nil {
		return nil, createClientInitializationError(operationAPIUpdate)
	}

	request := putClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(operationAPIUpdate)
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
		return createInvalidParameterError(operationAPIDelete, parameterSling)
	}

	if isEmpty(path) {
		return createInvalidParameterError(operationAPIDelete, parameterPath)
	}

	deleteClient := sling.Delete(path)
	if deleteClient == nil {
		return createClientInitializationError(operationAPIDelete)
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
