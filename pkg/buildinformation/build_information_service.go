package buildinformation

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	bulkPath string

	services.CanDeleteService
}

func NewBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *BuildInformationService {
	return &BuildInformationService{
		bulkPath: bulkPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceBuildInformationService, sling, uriTemplate),
		},
	}
}

func Add(client *http.Client, command *BuildInformationCommand) (*BuildInformationResponse, error) {
	// TODO: replace infrastructure (below)
	host := os.Getenv(constants.EnvironmentVariableOctopusHost)
	path, err := uritemplates.NewUriTemplateCache().Expand(uritemplates.BuildInformation, command)
	if err != nil {
		return nil, err
	}
	url := host + path

	buf, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	var data *BuildInformationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
