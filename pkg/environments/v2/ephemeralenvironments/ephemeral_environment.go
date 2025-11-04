package ephemeralenvironments

import (
	"math"
)

type EphemeralEnvironment struct {
	ID                  string `json:"Id"`
	Name                string `json:"Name"`
	SpaceID             string `json:"SpaceId"`
	Slug                string `json:"Slug"`
	Description         string `json:"Description"`
	Type                string `json:"Type"`
	SortOrder           int    `json:"SortOrder"`
	UseGuidedFailure    bool   `json:"UseGuidedFailure"`
	ParentEnvironmentId string `json:"ParentEnvironmentId"`
}

func NewEphemeralEnvironment(name string, parentEnvironmentID string, spaceID string) *EphemeralEnvironment {
	return &EphemeralEnvironment{
		Name:                name,
		SpaceID:             spaceID,
		SortOrder:           math.MaxInt,
		Type:                "Ephemeral",
		ParentEnvironmentId: parentEnvironmentID,
	}
}
