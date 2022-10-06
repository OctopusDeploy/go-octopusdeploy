package projects

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func createProjectService(t *testing.T) *ProjectService {
	service := NewProjectService(nil, constants.TestURIProjects, constants.TestURIProjectPulse, constants.TestURIProjectsExperimentalSummaries, constants.TestURIProjectsImportProjects, constants.TestURIProjectsExportProjects)
	services.NewServiceTests(t, service, constants.TestURIProjects, constants.ServiceProjectService)
	return service
}

func TestNewProjectService(t *testing.T) {
	ServiceFunction := NewProjectService
	client := &sling.Sling{}
	experimentalSummariesPath := ""
	pulsePath := ""
	importProjectsPath := ""
	exportProjectsPath := ""
	ServiceName := constants.ServiceProjectService
	uriTemplate := ""

	testCases := []struct {
		name                      string
		f                         func(*sling.Sling, string, string, string, string, string) *ProjectService
		client                    *sling.Sling
		uriTemplate               string
		experimentalSummariesPath string
		pulsePath                 string
		importProjectsPath        string
		exportProjectsPath        string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
		{"EmptyURITemplate", ServiceFunction, client, "", pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.pulsePath, tc.experimentalSummariesPath, tc.importProjectsPath, tc.exportProjectsPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectServiceGetWithEmptyID(t *testing.T) {
	service := createProjectService(t)

	resource, err := service.GetByID("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestProjectServiceGetByNameWithEmptyValue(t *testing.T) {
	service := createProjectService(t)

	resource, err := service.GetByName("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)
}
