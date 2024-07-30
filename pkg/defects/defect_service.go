package defects

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	template        = "/api/{spaceId}/releases/{releaseId}/defects"
	resolveTemplate = "/api/{spaceId}/releases/{releaseId}/defects/resolve"
)

// GetAll returns all defects for a release
func GetAll(client newclient.Client, spaceID string, releaseID string) ([]*Defect, error) {
	expandedUri, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId":   spaceID,
		"releaseId": releaseID,
	})
	if err != nil {
		return nil, err
	}

	return newclient.GetAll[Defect](client, expandedUri, spaceID)
}

// Create records a defect in a release
func Create(client newclient.Client, spaceID string, command *CreateReleaseDefectCommand) (*Defect, error) {
	expandedUri, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId":   spaceID,
		"releaseId": command.ReleaseID,
	})
	if err != nil {
		return nil, err
	}

	return newclient.Post[Defect](client.HttpSession(), expandedUri, command)
}

// Resolve resolves a defect in a release
func Resolve(client newclient.Client, spaceID string, command *ResolveReleaseDefectCommand) (*Defect, error) {
	expandedUri, err := client.URITemplateCache().Expand(resolveTemplate, map[string]any{
		"spaceId":   spaceID,
		"releaseId": command.ReleaseID,
	})
	if err != nil {
		return nil, err
	}

	return newclient.Post[Defect](client.HttpSession(), expandedUri, command)
}
