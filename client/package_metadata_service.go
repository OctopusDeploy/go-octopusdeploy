package client

import "github.com/dghubble/sling"

type packageMetadataService struct {
	service
}

func newPackageMetadataService(sling *sling.Sling, uriTemplate string) *packageMetadataService {
	return &packageMetadataService{
		service: newService(servicePackageMetadataService, sling, uriTemplate, nil),
	}
}
