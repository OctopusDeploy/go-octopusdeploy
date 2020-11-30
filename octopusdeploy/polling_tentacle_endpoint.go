package octopusdeploy

import (
	"encoding/json"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type PollingTentacleEndpoint struct {
	URI *url.URL `json:"Uri" validate:"required,uri"`

	tentacleEndpoint `validate:"required"`
}

func NewPollingTentacleEndpoint(uri *url.URL, thumbprint string) *PollingTentacleEndpoint {
	return &PollingTentacleEndpoint{
		tentacleEndpoint: *newTentacleEndpoint("TentacleActive", thumbprint),
		URI:              uri,
	}
}

func (l PollingTentacleEndpoint) MarshalJSON() ([]byte, error) {
	pollingTentacleEndpoint := struct {
		URI string `json:"Uri" validate:"required,uri"`
		tentacleEndpoint
	}{
		URI:              l.URI.String(),
		tentacleEndpoint: l.tentacleEndpoint,
	}

	return json.Marshal(pollingTentacleEndpoint)
}

// UnmarshalJSON sets this polling tentacle endpoint to its representation in
// JSON.
func (l *PollingTentacleEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		URI string `json:"Uri" validate:"required,uri"`
		tentacleEndpoint
	}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	err = validate.Struct(fields)
	if err != nil {
		return err
	}

	l.tentacleEndpoint = fields.tentacleEndpoint

	u, err := url.Parse(fields.URI)
	if err != nil {
		return err
	}
	l.URI = u

	return nil
}

// Validate checks the state of the listening tentacle endpoint and returns an
// error if invalid.
func (l *PollingTentacleEndpoint) Validate() error {
	return validator.New().Struct(l)
}
