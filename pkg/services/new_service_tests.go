package services

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/stretchr/testify/require"
)

func NewServiceTests(t *testing.T, service IService, uriTemplate string, ServiceName string) {
	require.NotNil(t, service)
	require.NotNil(t, service.GetClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.GetURITemplate(), template)
	require.Equal(t, service.GetName(), ServiceName)
}
