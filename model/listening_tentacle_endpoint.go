package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type ListeningTentacleEndpoint struct {
	ProxyID string  `json:"ProxyId"`
	URI     url.URL `json:"Uri" validate:"required,uri"`

	tentacleEndpoint
}

func NewListeningTentacleEndpoint(uri url.URL, thumbprint string) *ListeningTentacleEndpoint {
	resource := &ListeningTentacleEndpoint{}
	resource.CommunicationStyle = "TentaclePassive"
	resource.Thumbprint = thumbprint
	resource.URI = uri

	return resource
}

func (e ListeningTentacleEndpoint) MarshalJSON() ([]byte, error) {
	resource := struct {
		ProxyID string `json:"ProxyId"`
		URI     string `json:"Uri"`
		endpoint
	}{
		ProxyID:  e.ProxyID,
		URI:      e.URI.String(),
		endpoint: e.endpoint,
	}

	return json.Marshal(resource)
}

func (e *ListeningTentacleEndpoint) UnmarshalJSON(j []byte) error {
	s := strings.TrimSpace(string(j))

	var rawStrings map[string]interface{}
	err := json.Unmarshal([]byte(s), &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if k == "CommunicationStyle" {
			e.CommunicationStyle = v.(string)
		}
		if k == "Id" {
			if v != nil {
				e.ID = v.(string)
			}
		}
		if k == "LastModifiedOn" {
			if v != nil {
				e.LastModifiedOn = v.(*time.Time)
			}
		}
		if k == "LastModifiedBy" {
			if v != nil {
				e.LastModifiedBy = v.(string)
			}
		}
		if k == "Links" {
			links := v.(map[string]interface{})
			if len(links) > 0 {
				fmt.Println(links)
			}
		}
		if k == "ProxyId" {
			if v != nil {
				e.ProxyID = v.(string)
			}
		}
		if k == "URI" {
			rawURL := v.(string)
			u, err := url.Parse(rawURL)
			if err != nil {
				return err
			}
			e.URI = *u
		}
	}

	return nil
}

var _ ResourceInterface = &ListeningTentacleEndpoint{}
var _ EndpointInterface = &ListeningTentacleEndpoint{}
