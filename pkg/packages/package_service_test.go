package packages

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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

func TestParsePackageIdAndVersion(t *testing.T) {
	cannotParseError := errors.New("could not determine the package ID and/or version based on the supplied filename")

	testCases := []struct {
		fileName  string
		packageId string
		version   string
		err       error
	}{
		{"Octopus.Tentacle.6.3.417-x64.zip", "Octopus.Tentacle", "6.3.417-x64", nil},
		{"NuGet.CommandLine.6.2.3.nupkg", "NuGet.CommandLine", "6.2.3", nil},
		{"pterm.0.12.42.zip", "pterm", "0.12.42", nil},
		{"powershell-linux-alpine-x64.7.2.10.tar.gz", "powershell-linux-alpine-x64", "7.2.10", nil},
		{"powershell-linux-alpine-x64.7.2.10.tar.bz2", "powershell-linux-alpine-x64", "7.2.10", nil},
		{"powershell-linux-alpine-x64.7.2.10.tar", "powershell-linux-alpine-x64", "7.2.10", nil},

		// quirk: If someone zips an msi they may get .msi.zip. We remove the zip extension but .msi stays behind.
		// this matches the behaviour of the C# client
		{"Octopus.Tentacle.6.3.417-x64.msi.zip", "Octopus.Tentacle", "6.3.417-x64.msi", nil},

		// quirk. If someone uses a hyphen as a version separator then we think it's part of the package name.
		// this matches the C# client
		{"pterm-0.12.42.zip", "pterm-0", "12.42", nil},

		// error cases
		{"SqlServer2000.zip", "", "", cannotParseError},
		{"Octopus.Tentacle.zip", "", "", cannotParseError},
		{"Octopus.Tentacle-77.zip", "", "", cannotParseError},

		// case sensitivity
		{"Octopus.Tentacle.6.3.417-x64.ZIP", "Octopus.Tentacle", "6.3.417-x64", nil},
		{"NuGet.CommandLine.6.2.3.NUPKG", "NuGet.CommandLine", "6.2.3", nil},
		{"pterm.0.12.42.ZIP", "pterm", "0.12.42", nil},
		{"powershell-linux-alpine-x64.7.2.10.TAR.GZ", "powershell-linux-alpine-x64", "7.2.10", nil},
		{"powershell-linux-alpine-x64.7.2.10.TAR.BZ2", "powershell-linux-alpine-x64", "7.2.10", nil},
		{"powershell-linux-alpine-x64.7.2.10.TAR", "powershell-linux-alpine-x64", "7.2.10", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.fileName, func(t *testing.T) {
			packageId, version, err := ParsePackageIDAndVersion(tc.fileName)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.packageId, packageId)
			assert.Equal(t, tc.version, version)
		})
	}
}
