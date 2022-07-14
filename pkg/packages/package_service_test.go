package packages

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createPackageService(t *testing.T) *PackageService {
	service := NewPackageService(nil, constants.TestURIPackages, constants.TestURIPackageDeltaSignature, constants.TestURIPackageDeltaUpload, constants.TestURIPackageNotesList, constants.TestURIPackagesBulk, constants.TestURIPackageUpload)
	services.NewServiceTests(t, service, constants.TestURIPackages, constants.ServicePackageService)
	return service
}

func TestPackageServiceGetByID(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	require.Nil(t, resource)
}

func TestPackageServiceNew(t *testing.T) {
	ServiceFunction := NewPackageService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServicePackageService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string, string, string, string, string) *PackageService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, constants.TestURIPackageDeltaSignature, constants.TestURIPackageDeltaUpload, constants.TestURIPackageNotesList, constants.TestURIPackagesBulk, constants.TestURIPackageUpload)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestPackageServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createPackageService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID))
		})
	}
}

func TestPackageServiceUpdateWithEmptyPackage(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	updatedPackage, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, updatedPackage)

	updatedPackage, err = service.Update(&Package{})
	require.Error(t, err)
	require.Nil(t, updatedPackage)
}
