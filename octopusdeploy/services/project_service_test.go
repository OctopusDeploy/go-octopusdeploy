package services

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func createProjectService(t *testing.T) *projectService {
	service := newProjectService(nil, TestURIProjects, TestURIProjectPulse, TestURIProjectsExperimentalSummaries, TestURIProjectsImportProjects, TestURIProjectsExportProjects)
	testNewService(t, service, TestURIProjects, ServiceProjectService)
	return service
}

func TestNewProjectService(t *testing.T) {
	ServiceFunction := newProjectService
	client := &sling.Sling{}
	experimentalSummariesPath := emptyString
	pulsePath := emptyString
	importProjectsPath := emptyString
	exportProjectsPath := emptyString
	ServiceName := ServiceProjectService
	uriTemplate := emptyString

	testCases := []struct {
		name                      string
		f                         func(*sling.Sling, string, string, string, string, string) *projectService
		client                    *sling.Sling
		uriTemplate               string
		experimentalSummariesPath string
		pulsePath                 string
		importProjectsPath        string
		exportProjectsPath        string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, pulsePath, experimentalSummariesPath, importProjectsPath, exportProjectsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.pulsePath, tc.experimentalSummariesPath, tc.importProjectsPath, tc.exportProjectsPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectServiceGetWithEmptyID(t *testing.T) {
	service := createProjectService(t)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
