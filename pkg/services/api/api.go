package api

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/dghubble/sling"
)

var version = "development"
var UserAgentString = GetUserAgentString("")

// Generic OctopusDeploy API Get Function.
func ApiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIGet, "sling")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIGet)
	}

	client = client.Get(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIGet)
	}

	octopusDeployError := new(core.APIError)
	resp, err := client.Receive(inputStruct, &octopusDeployError)
	// if err != nil {
	// 	return nil, err
	// }

	// core.APIErrorChecker doesn't give us useful handling for auth errors; do this specifically
	if resp != nil && resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("unauthorized")
	}

	apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

func GetUserAgentString(requestingTool string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/OctopusDeploy/go-octopusdeploy/v2" {
				if dep.Version != "" {
					version = dep.Version
				}
			}
		}
	}

	automationEnvironment := GetAutomationEnvironment(&OsEnvironment{})

	return strings.TrimSpace(fmt.Sprintf("%s/%s (%s; %s) go/%s %s %s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version(), automationEnvironment, requestingTool))
}
