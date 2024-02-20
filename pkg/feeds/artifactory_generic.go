package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type ArtifactoryGenericFeed struct {
	LayoutRegex string `json:"LayoutRegex,omitempty"`
	FeedURI     string `json:"FeedUri,omitempty"`
	Repository  string `json:"Repository"`

	feed
}

func NewArtifactoryGenericFeed(name string) (*ArtifactoryGenericFeed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyError("name")
	}

	feed := ArtifactoryGenericFeed{
		feed: *newFeed(name, FeedTypeArtifactoryGeneric),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

func (g *ArtifactoryGenericFeed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(g)
}
