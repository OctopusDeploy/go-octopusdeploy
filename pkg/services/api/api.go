package api

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/dghubble/sling"
)

var version = "development"
var UserAgentString = GetUserAgentString()

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

	client.Set("User-Agent", UserAgentString)

	octopusDeployError := new(core.APIError)
	resp, err := client.Receive(inputStruct, &octopusDeployError)
	// if err != nil {
	// 	return nil, err
	// }

	apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

func GetUserAgentString() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/OctopusDeploy/go-octopusdeploy/v2" {
				if dep.Version != "" {
					version = dep.Version
				}
			}
		}
	}
	return fmt.Sprintf("%s/%s (%s; %s) go/%s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
