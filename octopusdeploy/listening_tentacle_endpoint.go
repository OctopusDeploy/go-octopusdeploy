package octopusdeploy

import (
	"encoding/json"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type ListeningTentacleEndpoint struct {
	ProxyID string   `json:"ProxyId,omitempty"`
	URI     *url.URL `json:"Uri" validate:"required,uri"`

	tentacleEndpoint
}

func NewListeningTentacleEndpoint(uri *url.URL, thumbprint string) *ListeningTentacleEndpoint {
	return &ListeningTentacleEndpoint{
		URI:              uri,
		tentacleEndpoint: *newTentacleEndpoint("TentaclePassive", thumbprint),
	}
}

func (l ListeningTentacleEndpoint) MarshalJSON() ([]byte, error) {
	listeningTentacleEndpoint := struct {
		ProxyID string `json:"ProxyId,omitempty"`
		URI     string `json:"Uri" validate:"required,uri"`
		tentacleEndpoint
	}{
		ProxyID:          l.ProxyID,
		URI:              l.URI.String(),
		tentacleEndpoint: l.tentacleEndpoint,
	}

	return json.Marshal(listeningTentacleEndpoint)
}

// UnmarshalJSON sets this listening tentacle endpoint to its representation in
// JSON.
func (l *ListeningTentacleEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		ProxyID string `json:"ProxyId,omitempty"`
		URI     string `json:"Uri" validate:"required,uri"`
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

	l.ProxyID = fields.ProxyID
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
func (l *ListeningTentacleEndpoint) Validate() error {
	return validator.New().Struct(l)
}
