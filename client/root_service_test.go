package client

import (
	"testing"

	"github.com/dghubble/sling"
)

func TestNewRootService(t *testing.T) {
	serviceFunction := newRootService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceRootService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *rootService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}
